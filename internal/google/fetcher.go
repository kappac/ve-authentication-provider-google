package google

import (
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/micro/micro/v3/service/logger"
)

const (
	fetcherGetTimeout            = 10 * time.Second
	fetcherCertsURL              = "https://www.googleapis.com/oauth2/v1/certs"
	fetcherCacheControlHeaderKey = "Cache-Control"
)

type fetcherCertsMap map[string][]byte

type fetcherProcessingResult struct {
	maxAge time.Duration
	certs  fetcherCertsMap
	err    error
}

type fetcherGetResult struct {
	resp http.Response
	err  error
}

type fetcher struct {
	closeCh      chan chan error
	fetchGetCh   <-chan fetcherGetResult
	processingCh chan fetcherProcessingResult
	updatesCh    chan fetcherCertsMap
	nextUpdate   time.Time
	nextUpdateCh <-chan time.Time
	isClosing    bool
	err          error
}

func newFetcher() *fetcher {
	return &fetcher{
		closeCh:   make(chan chan error),
		updatesCh: make(chan fetcherCertsMap),
	}
}

func (f *fetcher) run() {
	for !f.isClosing {
		select {
		case errc := <-f.closeCh:
			logger.Debug("Closing")

			f.isClosing = true
			f.closeChannels()
			errc <- f.err
		case <-f.nextUpdateCh:
			logger.Debug("next update: request")

			if f.fetchGetCh == nil {
				logger.Debug("next update: fetch")

				f.fetchGetCh = f.get()
			}
		case gr := <-f.fetchGetCh:
			f.fetchGetCh = nil

			logger.Debugf("fetch result: %v\n", gr)

			if gr.err != nil {
				f.err = gr.err
				f.setNextUpdate(fetcherGetTimeout)
				f.schedule()
				break
			}

			f.err = nil

			if f.processingCh == nil {
				f.processingCh = f.processResponse(gr.resp)
			}
		case pr := <-f.processingCh:
			logger.Debugf("processing response: %v\n", pr)

			f.processingCh = nil

			var (
				nt time.Duration
			)

			if pr.err != nil {
				f.err = pr.err
				nt = fetcherGetTimeout
			} else {
				f.updatesCh <- pr.certs

				f.err = nil
				nt = pr.maxAge
			}

			f.setNextUpdate(nt)
			f.schedule()
		}
	}
}

func (f *fetcher) stop() error {
	cc := make(chan error)

	f.closeCh <- cc

	select {
	case errc := <-cc:
		return errc
	}
}

func (f *fetcher) subscribe() chan fetcherCertsMap {
	return f.updatesCh
}

func (f *fetcher) closeChannels() {
	close(f.closeCh)
	close(f.updatesCh)
}

func (f *fetcher) setNextUpdate(d time.Duration) {
	f.nextUpdate = time.Now().Add(d)
}

func (f *fetcher) schedule() {
	if f.nextUpdate.IsZero() {
		f.nextUpdate = time.Now()
	}

	d := f.nextUpdate.Sub(time.Now())

	f.nextUpdateCh = time.After(d)
}

func (f *fetcher) get() <-chan fetcherGetResult {
	c := make(chan fetcherGetResult)

	go func() {
		resp, err := http.Get(fetcherCertsURL)
		c <- fetcherGetResult{
			resp: *resp,
			err:  err,
		}
	}()

	return c
}

func (f *fetcher) processResponse(r http.Response) chan fetcherProcessingResult {
	c := make(chan fetcherProcessingResult)

	go func() {
		maxAge, err := getMaxAge(&r.Header)
		if err != nil {
			c <- fetcherProcessingResult{
				err: err,
			}

			logger.Fatal(err)

			return
		}

		certs, err := getCerts(&r.Body)
		if err != nil {
			c <- fetcherProcessingResult{
				err: err,
			}

			logger.Fatal(err)

			return
		}

		c <- fetcherProcessingResult{
			maxAge: maxAge,
			certs:  certs,
		}
	}()

	return c
}

func getMaxAge(h *http.Header) (time.Duration, error) {
	cacheControlHeader := h.Get(fetcherCacheControlHeaderKey)

	if cacheControlHeader == "" {
		errStr := fmt.Sprintf("\"%s\" is abscent in the response header", fetcherCacheControlHeaderKey)
		return 0, errors.New(errStr)
	}

	re := regexp.MustCompile(`max-age=(\d*)`)

	res := re.FindStringSubmatch(cacheControlHeader)
	if res == nil {
		errStr := fmt.Sprintf("\"max-age\" property is absent in \"%s\" key", fetcherCacheControlHeaderKey)
		return 0, errors.New(errStr)
	}

	resLen := len(res)
	if resLen <= 1 {
		errStr := "The value for \"max-age\" property is absent"
		return 0, errors.New(errStr)
	}

	nt, err := strconv.Atoi(res[resLen-1])
	if err != nil {
		return 0, err
	}

	return time.Duration(nt) * time.Second, nil
}

func getCerts(r *io.ReadCloser) (fetcherCertsMap, error) {
	certs := make(fetcherCertsMap)

	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return nil, err
	}

	p := new(map[string][]byte)
	if err := json.Unmarshal(b, p); err != nil {
		return nil, err
	}

	logger.Debugf("JSON from %s: %v", fetcherCertsURL, p)

	var decodeErr error

	for k, v := range *p {
		b, _ := pem.Decode(v)
		if b == nil {
			errStr := fmt.Sprintf("No certificate for %s", k)
			decodeErr = errors.New(errStr)
			break
		}
	}

	return certs, decodeErr
}

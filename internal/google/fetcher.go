package google

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

type fetcherCertsMap map[string]string

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
	logger       logger.Logger
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
		logger: logger.New(
			logger.WithEntity("Fetcher"),
		),
		closeCh:   make(chan chan error),
		updatesCh: make(chan fetcherCertsMap),
	}
}

func (f *fetcher) run() {
	f.logger.Debugm("starting")

	f.schedule()

	for !f.isClosing {
		select {
		case errc := <-f.closeCh:
			f.isClosing = true
			f.closeChannels()

			f.logger.Debugm("closing", "err", f.err)

			errc <- f.err
		case <-f.nextUpdateCh:
			if f.fetchGetCh == nil {
				f.logger.Debugm("new request")

				f.fetchGetCh = f.get()
			}
		case gr := <-f.fetchGetCh:
			f.fetchGetCh = nil

			f.logger.Infom("request finished", "err", gr.err)

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
			f.processingCh = nil

			var (
				nt time.Duration
			)

			f.logger.Infom("response processed", "err", pr.err)

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
	f.logger.Debugm("stopping")

	cc := make(chan error)
	f.closeCh <- cc
	return <-cc
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

			return
		}

		certs, err := getCerts(r.Body)
		if err != nil {
			c <- fetcherProcessingResult{
				err: err,
			}

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
		return 0, errorCacheControlAbsent
	}

	re := regexp.MustCompile(`max-age=(\d*)`)

	res := re.FindStringSubmatch(cacheControlHeader)
	if res == nil {
		return 0, errorMaxAgePropertyAbsent
	}

	resLen := len(res)
	if resLen <= 1 {
		return 0, errorMaxAgeValueAbsent
	}

	nt, err := strconv.Atoi(res[resLen-1])
	if err != nil {
		return 0, errorMaxAgeConvertion
	}

	return time.Duration(nt) * time.Second, nil
}

func getCerts(r io.ReadCloser) (fetcherCertsMap, error) {
	defer r.Close()

	certs := make(fetcherCertsMap)
	if err := json.NewDecoder(r).Decode(&certs); err != nil {
		return nil, errorJSONDecode
	}

	return certs, nil
}

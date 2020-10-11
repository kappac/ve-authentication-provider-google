package google

import (
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/micro/micro/v3/service/logger"
)

type oauthCertificatesMap map[string]*pem.Block

type oauthCertificates struct {
	f             *fetcher
	certsMap      oauthCertificatesMap
	certsUpdateCh chan fetcherCertsMap
	closeCh       chan chan error
	isClosing     bool
	err           error
}

func newOauthCertificates() *oauthCertificates {
	f := newFetcher()

	certsUpdateCh := f.subscribe()
	closeCh := make(chan chan error)

	go f.run()

	return &oauthCertificates{
		f:             f,
		certsUpdateCh: certsUpdateCh,
		closeCh:       closeCh,
	}
}

func (oc *oauthCertificates) run() {
	for !oc.isClosing {
		select {
		case errc := <-oc.closeCh:
			logger.Debug("Closing")
			oc.isClosing = true

			oc.err = oc.f.stop()

			oc.closeChannels()
			errc <- oc.err
		case certsMap := <-oc.certsUpdateCh:
			logger.Debugf("Certifications map update: %v", certsMap)
			oc.err = oc.processCertsMap(certsMap)

			if oc.err != nil {
				logger.Fatal(oc.err)
			}
		}
	}
}

func (oc *oauthCertificates) stop() error {
	cc := make(chan error)

	oc.closeCh <- cc

	select {
	case errc := <-cc:
		return errc
	}
}

func (oc *oauthCertificates) closeChannels() {
	close(oc.closeCh)
}

func (oc *oauthCertificates) processCertsMap(cm fetcherCertsMap) error {
	var decodeErr error
	oc.certsMap = make(oauthCertificatesMap)

	for k, v := range cm {
		b, _ := pem.Decode(v)

		if b == nil {
			errStr := fmt.Sprintf("No certificate for %s", k)
			decodeErr = errors.New(errStr)
			break
		}

		oc.certsMap[k] = b
	}

	if decodeErr != nil {
		return decodeErr
	}

	return nil
}

func (oc *oauthCertificates) get(k string) (*pem.Block, error) {
	var (
		b  *pem.Block
		ok bool
	)

	if b, ok = oc.certsMap[k]; !ok {
		errStr := fmt.Sprintf("No certificate for %s", k)
		err := errors.New(errStr)

		logger.Error(errStr)

		return nil, err
	}

	return b, nil
}

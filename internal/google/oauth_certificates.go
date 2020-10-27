package google

import (
	"encoding/pem"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

type oauthCertificatesMap map[string]*pem.Block

type oauthCertificates struct {
	logger        logger.Logger
	f             *fetcher
	certsMap      oauthCertificatesMap
	certsUpdateCh chan fetcherCertsMap
	closeCh       chan chan error
	isClosing     bool
	err           error
}

func newOauthCertificates() *oauthCertificates {
	f := newFetcher()

	l := logger.New(
		logger.WithEntity("OAuthCertificates"),
	)

	certsUpdateCh := f.subscribe()
	closeCh := make(chan chan error)

	go f.run()

	return &oauthCertificates{
		logger:        l,
		f:             f,
		certsUpdateCh: certsUpdateCh,
		closeCh:       closeCh,
	}
}

func (oc *oauthCertificates) run() {
	oc.logger.Debugm("starting")

	for !oc.isClosing {
		select {
		case errc := <-oc.closeCh:
			oc.isClosing = true

			oc.err = oc.f.stop()

			oc.logger.Debugm("closing", "err", oc.err)

			oc.closeChannels()
			errc <- oc.err
		case certsMap := <-oc.certsUpdateCh:
			oc.err = oc.processCertsMap(certsMap)

			oc.logger.Infom("certificates processed", "err", oc.err)
		}
	}
}

func (oc *oauthCertificates) stop() error {
	oc.logger.Debugm("stopping")

	cc := make(chan error)
	oc.closeCh <- cc
	return <-cc
}

func (oc *oauthCertificates) closeChannels() {
	close(oc.closeCh)
}

func (oc *oauthCertificates) processCertsMap(cm fetcherCertsMap) error {
	var decodeErr error
	oc.certsMap = make(oauthCertificatesMap)

	for k, v := range cm {
		b, _ := pem.Decode([]byte(v))

		if b == nil {
			decodeErr = errorNoCertificate
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
		return nil, errorNoCertificate
	}

	return b, nil
}

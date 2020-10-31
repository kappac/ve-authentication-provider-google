package google

import (
	"encoding/pem"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

type oauthCertificatesMap map[string]*pem.Block

type oauthCertificatesStatisticsUpdate struct {
	keys []string
	err  error
}

type oauthCertificates struct {
	logger        logger.Logger
	f             *fetcher
	certsMap      oauthCertificatesMap
	certsUpdateCh chan fetcherUpdateCerts
	statisticsCh  chan oauthCertificatesStatisticsUpdate
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
	statisticsCh := make(chan oauthCertificatesStatisticsUpdate)
	closeCh := make(chan chan error)

	go f.run()

	return &oauthCertificates{
		logger:        l,
		f:             f,
		certsUpdateCh: certsUpdateCh,
		closeCh:       closeCh,
		statisticsCh:  statisticsCh,
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

			errc <- oc.err
			oc.closeChannels()
		case certsUpdate, ok := <-oc.certsUpdateCh:
			if !ok {
				break
			}

			oc.err = certsUpdate.err

			if oc.err == nil {
				oc.err = oc.processCertsMap(certsUpdate.certs)
			}

			oc.statisticsCh <- oauthCertificatesStatisticsUpdate{
				keys: oc.getKeys(),
				err:  oc.err,
			}

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
	close(oc.statisticsCh)
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

func (oc *oauthCertificates) subscribeStatistics() <-chan oauthCertificatesStatisticsUpdate {
	return oc.statisticsCh
}

func (oc *oauthCertificates) getKeys() []string {
	var r []string

	for k := range oc.certsMap {
		r = append(r, k)
	}

	return r
}

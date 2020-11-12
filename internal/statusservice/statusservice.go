package statusservice

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"sync"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/veservice"
)

// Probe type
type Probe string

// Probe types supported by the Service
const (
	ProbeLiveness  Probe = "probe_liveness"
	ProbeReadiness       = "probe_readiness"
	ProbeStartup         = "probe_startup"
)

const (
	endpointURI          = "/probes"
	endpointURILiveness  = endpointURI + "/liveness"
	endpointURIReadiness = endpointURI + "/readiness"
	endpointURIStartup   = endpointURI + "/startup"
)

var (
	probeEndpoint = map[Probe]string{
		ProbeLiveness:  endpointURILiveness,
		ProbeReadiness: endpointURIReadiness,
		ProbeStartup:   endpointURIStartup,
	}
)

// SourceData describes a piece of data operated by SourceSubscriber
type SourceData interface {
	GetError() error
	GetData() map[string]interface{}
}

// SourceSubscriber ...
type SourceSubscriber interface {
	Subscribe() <-chan SourceData
	Unsubscribe() error
}

// StatusService starts an http server with health check endpoints.
type StatusService interface {
	veservice.RunStopper
}

type closing chan interface{}

type sources map[Probe]SourceSubscriber

type subscribes []reflect.SelectCase

type probeMap map[int]Probe

type probeData map[Probe]SourceData

type probeEndpoints map[Probe]endpoint.Endpoint

type statusService struct {
	wg           sync.WaitGroup
	once         sync.Once
	runOnce      sync.Once
	logger       logger.Logger
	address      string
	server       *http.Server
	sources      sources
	subscribes   subscribes
	subsProbeMap probeMap
	datas        probeData
	datasMX      sync.Mutex
	endpoints    probeEndpoints
	closing      closing
	err          error
}

// New constructs new instance of a Service
func New(opts ...NewOption) StatusService {
	s := &statusService{
		logger: logger.New(
			logger.WithEntity("StatusService"),
		),
		sources: make(sources),
	}

	for _, o := range opts {
		o(s)
	}

	s.server = &http.Server{
		Addr: s.address,
	}

	sourcesAmount := len(s.sources)
	// the last subscribe is closing chan.
	subscribesAmount := sourcesAmount + 1
	s.closing = make(closing)
	s.subscribes = make(subscribes, subscribesAmount)
	s.subsProbeMap = make(probeMap, sourcesAmount)
	s.datas = make(probeData, sourcesAmount)
	s.endpoints = make(probeEndpoints, sourcesAmount)

	s.subscribeSources()
	s.generateEndpoints()
	go s.listenSubscribes()

	return s
}

func (s *statusService) Run() error {
	s.logger.Debugm("Run")

	if len(s.sources) == 0 {
		return errorNoProbesConfigured
	}

	return s.runServer()
}

func (s *statusService) Stop() error {
	s.logger.Debugm("Stop")

	s.once.Do(func() {
		close(s.closing)
	})
	s.wg.Wait()

	return s.err
}

func (s *statusService) stop() {
	s.unsubscribeSources()
	s.stopServer()
}

func (s *statusService) subscribeSources() {
	s.logger.Debugm("subscribeSource")

	s.wg.Add(1)

	for k, v := range s.sources {
		nextCaseIndex := len(s.subsProbeMap)
		s.subscribes[nextCaseIndex] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(v.Subscribe()),
		}
		s.subsProbeMap[nextCaseIndex] = k
	}

	cidx := len(s.subscribes) - 1

	s.subscribes[cidx] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(s.closing),
	}
}

func (s *statusService) unsubscribeSources() {
	s.logger.Debugm("unsubscribeSources")

	defer s.wg.Done()

	for _, src := range s.sources {
		s.err = src.Unsubscribe()
	}

	sc := len(s.subscribes)

	for idx, sbs := range s.subscribes {
		if idx < sc {
			sbs.Chan = reflect.ValueOf(nil)
		}
	}
}

func (s *statusService) listenSubscribes() {
	s.logger.Debugm("ListenSubscribes")

	// the index of the closing channel
	cidx := len(s.subscribes) - 1

	for {
		cID, v, ok := reflect.Select(s.subscribes)
		if cID == cidx {
			s.stop()
			return
		}

		if !ok {
			s.subscribes[cID].Chan = reflect.ValueOf(nil)
		} else {
			s.processChannelRecv(
				s.subsProbeMap[cID],
				v.Interface().(SourceData),
			)
		}
	}
}

func (s *statusService) processChannelRecv(p Probe, sd SourceData) {
	s.datasMX.Lock()
	defer s.datasMX.Unlock()
	s.logger.Debugm("processChannelRecv", "probe", p)
	s.datas[p] = sd
}

func (s *statusService) generateEndpoints() {
	for p := range s.sources {
		s.logger.Debugm("generateEndpoints", "probe", p)
		s.endpoints[p] = s.newEndpoint(p)
	}
}

func (s *statusService) runServer() error {
	s.logger.Debugm("runServer")

	handler := http.NewServeMux()

	for p, e := range s.endpoints {
		if uri, ok := probeEndpoint[p]; ok {
			s.logger.Debugm("registeringEndpoint", "probe", p, "endpoint", uri)

			handler.Handle(
				uri,
				httptransport.NewServer(e, decodeRequest, encodeResponse),
			)
		}
	}

	s.server.Handler = handler

	s.wg.Add(1)

	return s.server.ListenAndServe()
}

func (s *statusService) stopServer() {
	s.logger.Debugm("stopServer")
	s.err = s.server.Close()
	s.wg.Done()
}

func (s *statusService) newEndpoint(p Probe) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		s.datasMX.Lock()
		defer s.datasMX.Unlock()
		sd, ok := s.datas[p]
		if !ok {
			return nil, errorNoSourceData
		}

		return sd.GetData(), sd.GetError()
	}
}

func decodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

package statusservice

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
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
	RunStopper
}

type statusService struct {
	logger       logger.Logger
	address      string
	server       *http.Server
	sources      map[Probe]SourceSubscriber
	subscribes   []reflect.SelectCase
	subsProbeMap map[int]Probe
	datas        map[Probe]SourceData
	endpoints    map[Probe]endpoint.Endpoint
	isClosing    bool
}

// New constructs new instance of a Service
func New(opts ...NewOption) StatusService {
	s := &statusService{
		logger: logger.New(
			logger.WithEntity("StatusService"),
		),
		sources: make(map[Probe]SourceSubscriber),
	}

	for _, o := range opts {
		o(s)
	}

	s.server = &http.Server{
		Addr: s.address,
	}

	sourcesAmount := len(s.sources)
	s.subscribes = make([]reflect.SelectCase, sourcesAmount)
	s.subsProbeMap = make(map[int]Probe, sourcesAmount)
	s.datas = make(map[Probe]SourceData, sourcesAmount)
	s.endpoints = make(map[Probe]endpoint.Endpoint, sourcesAmount)

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

	s.isClosing = true
	s.unsubscribeSources()
	return s.stopServer()
}

func (s *statusService) subscribeSources() {
	s.logger.Debugm("subscribeSource")

	for k, v := range s.sources {
		nextCaseIndex := len(s.subsProbeMap)
		s.subscribes[nextCaseIndex] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(v.Subscribe()),
		}
		s.subsProbeMap[nextCaseIndex] = k
	}
}

func (s *statusService) unsubscribeSources() error {
	s.logger.Debugm("unsubscribeSources")

	var err error

	for _, s := range s.sources {
		err = s.Unsubscribe()
	}

	return err
}

func (s *statusService) listenSubscribes() {
	s.logger.Debugm("ListenSubscribes")

	for !s.isClosing {
		cID, v, ok := reflect.Select(s.subscribes)
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

	return s.server.ListenAndServe()
}

func (s *statusService) stopServer() error {
	s.logger.Debugm("stopServer")
	return s.server.Close()
}

func (s *statusService) newEndpoint(p Probe) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
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

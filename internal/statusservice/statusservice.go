package statusservice

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kappac/ve-authentication-provider-google/internal/types/runstopper"
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
	Subscribe() chan<- SourceData
}

// Service starts an http server with health check endpoints.
type Service interface {
	runstopper.RunStopper
}

type service struct {
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
func New(opts ...NewOption) Service {
	s := &service{
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

func (s *service) Run() error {
	if len(s.sources) == 0 {
		return errorNoProbesConfigured
	}

	return s.runServer()
}

func (s *service) Stop() error {
	s.isClosing = true
	return s.stopServer()
}

func (s *service) subscribeSources() {
	for k, v := range s.sources {
		nextCaseIndex := len(s.subsProbeMap)
		s.subscribes[nextCaseIndex] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(v.Subscribe()),
		}
		s.subsProbeMap[nextCaseIndex] = k
	}
}

func (s *service) listenSubscribes() {
	for !s.isClosing {
		cID, v, ok := reflect.Select(s.subscribes)
		if !ok {
			s.subscribes[cID].Chan = reflect.ValueOf(nil)
		}
		s.processChannelRecv(
			s.subsProbeMap[cID],
			v.Interface().(SourceData),
		)
	}
}

func (s *service) processChannelRecv(p Probe, sd SourceData) {
	s.datas[p] = sd
}

func (s *service) generateEndpoints() {
	for k, v := range s.datas {
		s.endpoints[k] = newEndpoint(v)
	}
}

func (s *service) runServer() error {
	handler := http.NewServeMux()

	for p, e := range s.endpoints {
		if uri, ok := probeEndpoint[p]; ok {
			handler.Handle(
				uri,
				httptransport.NewServer(e, decodeRequest, encodeResponse),
			)
		}
	}

	s.server.Handler = handler

	return s.server.ListenAndServe()
}

func (s *service) stopServer() error {
	return s.server.Close()
}

func newEndpoint(sd SourceData) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return sd.GetData(), sd.GetError()
	}
}

func decodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

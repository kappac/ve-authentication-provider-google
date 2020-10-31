package google

import "github.com/kappac/ve-authentication-provider-google/internal/statusservice"

// StatisticsSourceData is suitable for StatusService processing
type StatisticsSourceData struct {
	err  error
	data map[string]interface{}
}

// GetError returns an error for this peace of statistics.
func (d StatisticsSourceData) GetError() error {
	return d.err
}

// GetData returns a data for this peace of statistics.
func (d StatisticsSourceData) GetData() map[string]interface{} {
	return d.data
}

// StatisticsSource is intended to provide source data for StatusService
type StatisticsSource interface {
	statusservice.SourceSubscriber
}

type statisticsSource struct {
	oauthSubscribe <-chan oauthCertificatesStatisticsUpdate
	updatec        chan statusservice.SourceData
	isClosing      bool
}

// NewStatisticsSource constructs new StatisticsSource instance
// to be used with StatusService
func NewStatisticsSource(opts ...StatisticsSourceOption) StatisticsSource {
	s := &statisticsSource{
		updatec: make(chan statusservice.SourceData, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	go s.start()

	return s
}

func (s *statisticsSource) Subscribe() <-chan statusservice.SourceData {
	return s.updatec
}

func (s *statisticsSource) Unsubscribe() error {
	s.isClosing = true
	close(s.updatec)
	return nil
}

func (s *statisticsSource) start() {
	for !s.isClosing {
		select {
		case stats, ok := <-s.oauthSubscribe:
			if ok {
				s.updatec <- s.constructSourceData(stats)
			}
		}
	}
}

func (s *statisticsSource) constructSourceData(u oauthCertificatesStatisticsUpdate) StatisticsSourceData {
	var data = make(map[string]interface{})

	data["keys"] = u.keys

	return StatisticsSourceData{
		err:  u.err,
		data: data,
	}
}

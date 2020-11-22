package google

import (
	"sync"

	"github.com/kappac/ve-back-end-utils/pkg/statusservice"
)

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

type closing chan interface{}

type statisticsSource struct {
	wg             sync.WaitGroup
	once           sync.Once
	oauthSubscribe <-chan oauthCertificatesStatisticsUpdate
	updatec        chan statusservice.SourceData
	closing        closing
}

// NewStatisticsSource constructs new StatisticsSource instance
// to be used with StatusService
func NewStatisticsSource(opts ...StatisticsSourceOption) StatisticsSource {
	s := &statisticsSource{
		updatec: make(chan statusservice.SourceData),
		closing: make(closing),
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
	s.once.Do(func() {
		close(s.closing)
	})
	s.wg.Wait()
	return nil
}

func (s *statisticsSource) start() {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case stats, ok := <-s.oauthSubscribe:
			if ok {
				s.updatec <- s.constructSourceData(stats)
			}
		case <-s.closing:
			close(s.updatec)
			return
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

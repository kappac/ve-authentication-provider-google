package google

// StatisticsSourceOption ...
type StatisticsSourceOption func(s *statisticsSource)

// WithTokenVerifier sets a TokenVerifier to watch on.
func WithTokenVerifier(tv TokenVerifier) StatisticsSourceOption {
	return func(s *statisticsSource) {
		s.oauthSubscribe = tv.(*tokenVerifier).certs.subscribeStatistics()
	}
}

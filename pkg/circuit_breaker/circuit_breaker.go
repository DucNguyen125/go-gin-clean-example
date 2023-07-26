package circuit_breaker

import (
	"github.com/sony/gobreaker"
)

type Service interface {
	Execute(req func() (interface{}, error)) (interface{}, error)
}

type circuitBreakerService struct {
	*gobreaker.CircuitBreaker
}

func NewCircuitBreakerService(name string) Service {
	var st gobreaker.Settings
	st.Name = name
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	cb := gobreaker.NewCircuitBreaker(st)
	return &circuitBreakerService{cb}
}

func (s *circuitBreakerService) Execute(req func() (interface{}, error)) (interface{}, error) {
	return s.Execute(req) //nolint:staticcheck // no-recursive
}

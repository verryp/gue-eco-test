package service

import "context"

type IHealthCheck interface {
	HealthCheck(ctx context.Context) error
}

type healthCheck struct {
	*Option
}

func NewHealthCheckService(opt *Option) IHealthCheck {
	return &healthCheck{opt}
}

func (s *healthCheck) HealthCheck(_ context.Context) error {
	return nil
}

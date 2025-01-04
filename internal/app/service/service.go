package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {

	data := time.Date(2025, time.May, 13, 18, 40, 0, 0, time.UTC)

	dur := time.Until(data)

	return int64(dur.Hours()) / 24
}

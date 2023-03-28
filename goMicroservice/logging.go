package main

import (
	"context"
	"log"
	"time"
)

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		log.Printf("took: %v\nprice: %v\nerr: %v", time.Since(begin), price, err)
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

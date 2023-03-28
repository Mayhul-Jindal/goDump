package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface{
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetch(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 22020.4,
	"ETH": 22020.4,
	"HNN": 22020.4,
}

func MockPriceFetch(ctx context.Context, ticker string)(float64, error){
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
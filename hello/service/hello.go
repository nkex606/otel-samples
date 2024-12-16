package service

import (
	"context"
	"fmt"
	"otel-hello/domain"
)

type helloService struct {
	worldMicroSvc domain.WorldMicroSvc
}

func NewHelloService(wmsvc domain.WorldMicroSvc) *helloService {
	return &helloService{
		worldMicroSvc: wmsvc,
	}
}

func (h *helloService) CallWorld(ctx context.Context, cityId string) (string, error) {
	city, err := h.worldMicroSvc.CallWorldServer(ctx, cityId)
	if err != nil {
		fmt.Printf("[Err] service CallWorld fail: %s", err.Error())
		return "", err
	}
	return city, nil
}

func (h *helloService) CallWorldWithCapital(ctx context.Context, cityName string) (string, error) {
	capital, err := h.worldMicroSvc.CallWorldServerWithCapital(ctx, cityName)
	if err != nil {
		fmt.Printf("[Err] service CallWorldWithCapital fail: %s", err.Error())
		return "", err
	}
	return capital, nil
}

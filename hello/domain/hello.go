package domain

import "context"

type Helloworld struct {
	Name string `json:"name"`
}

type HelloService interface {
	CallWorld(ctx context.Context, cityId string) (string, error)
	CallWorldWithCapital(ctx context.Context, cityName string) (string, error)
}

type WorldMicroSvc interface {
	CallWorldServer(ctx context.Context, cityId string) (string, error)
	CallWorldServerWithCapital(ctx context.Context, cityName string) (string, error)
}

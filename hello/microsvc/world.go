package microservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"otel-hello/config"
	"otel-hello/domain"
)

type worldMicroSvc struct {
	client http.Client
}

func NewWorldMicroSvc(c http.Client) *worldMicroSvc {
	return &worldMicroSvc{
		client: c,
	}
}

func (w *worldMicroSvc) CallWorldServer(ctx context.Context, cityId string) (string, error) {
	worldServerApi := fmt.Sprintf("%s/%s/%s", config.GetWorldServerConfig().Host, "city", cityId)

	req, err := http.NewRequestWithContext(ctx, "GET", worldServerApi, nil)
	if err != nil {
		fmt.Printf("[Err] fail to new http request: %s", err.Error())
		return "", err
	}
	resp, err := w.client.Do(req)
	if err != nil {
		fmt.Printf("[Err] fail to do request: %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	city := domain.Helloworld{}
	if err = json.Unmarshal(data, &city); err != nil {
		fmt.Printf("[Err] fail to unmarshal data: %s", err.Error())
	}
	return city.Name, nil
}

func (w *worldMicroSvc) CallWorldServerWithCapital(ctx context.Context, cityName string) (string, error) {
	worldServerApi := fmt.Sprintf("%s/%s/%s", config.GetWorldServerConfig().Host, "city/capital", cityName)
	req, err := http.NewRequestWithContext(ctx, "GET", worldServerApi, nil)
	if err != nil {
		fmt.Printf("[Err] fail to new http request: %s", err.Error())
		return "", err
	}
	resp, err := w.client.Do(req)
	if err != nil {
		fmt.Printf("[Err] fail to do request: %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	capital := domain.Helloworld{}
	if err = json.Unmarshal(data, &capital); err != nil {
		fmt.Printf("[Err] fail to unmarshal data: %s", err.Error())
	}

	return capital.Name, nil
}

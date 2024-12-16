package service

import (
	"context"
	"fmt"
	"otel-world/model"
)

type cityService struct {
	cityRepo model.CityRepo
}

func NewCityService(cr model.CityRepo) *cityService {
	return &cityService{
		cityRepo: cr,
	}
}

func (cs *cityService) CityNameById(ctx context.Context, id int) (string, error) {
	city, err := cs.cityRepo.GetCityById(ctx, id)
	if err != nil {
		fmt.Printf("[Err] city service CityNameById: %s\n", err.Error())
		return "", err
	}
	return city.Name, nil
}

func (cs *cityService) CapitalNameByCity(ctx context.Context, cityName string) (string, error) {
	city, err := cs.cityRepo.GetCapitalByCity(ctx, cityName)
	if err != nil {
		fmt.Printf("[Err] city service CapitalNameByCity: %s\n", err.Error())
		return "", err
	}
	return city.Capital, nil
}

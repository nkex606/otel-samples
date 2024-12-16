package model

import "context"

type City struct {
	Name    string `json:"name" gorm:"column:name"`
	Capital string `json:"capital" gorm:"column:capital"`
}

func (c City) TableName() string {
	return "citys"
}

type CityRepo interface {
	GetCityById(ctx context.Context, id int) (City, error)
	GetCapitalByCity(ctx context.Context, cityName string) (City, error)
}

type CityService interface {
	CityNameById(ctx context.Context, id int) (string, error)
	CapitalNameByCity(ctx context.Context, cityName string) (string, error)
}

package repo

import (
	"context"
	"errors"
	"otel-world/model"

	"gorm.io/gorm"
)

type cityRepo struct {
	db *gorm.DB
}

func NewCityRepo(db *gorm.DB) *cityRepo {
	return &cityRepo{
		db: db,
	}
}

func (c *cityRepo) GetCityById(ctx context.Context, id int) (city model.City, err error) {
	city = model.City{}
	result := c.db.WithContext(ctx).Where("id = ?", id).Take(&city)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return city, result.Error
	}
	return city, nil
}

func (c *cityRepo) GetCapitalByCity(ctx context.Context, cityName string) (city model.City, err error) {
	city = model.City{}
	result := c.db.WithContext(ctx).Where("name = ?", cityName).Take(&city)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return city, result.Error
	}
	return city, nil
}

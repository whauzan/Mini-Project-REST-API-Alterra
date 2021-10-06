package food

import (
	"miniproject/business/food"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Name        string
	Photo       string
	Summary     string
	Number      int
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
}

func toDomain(record Food) food.Domain {
	return food.Domain{
		ID:          record.ID,
		Name:        record.Name,
		Photo:       record.Photo,
		Summary:     record.Summary,
		Number:      record.Number,
		Step:        record.Step,
		HealthScore: record.HealthScore,
		DishTypes:   record.DishTypes,
		Diets:       record.Diets,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func fromDomain(record Food) food.Domain {
	return Food{
		ID:          record.ID,
		Name:        record.Name,
		Photo:       record.Photo,
		Summary:     record.Summary,
		Number:      record.Number,
		Step:        record.Step,
		HealthScore: record.HealthScore,
		DishTypes:   record.DishTypes,
		Diets:       record.Diets,
	}
}

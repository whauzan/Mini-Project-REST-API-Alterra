package response

import (
	"miniproject/business/food"
	"time"
)

type Food struct {
	ID          int       `json:"id"`
	Photo       string    `json:"photo"`
	Summary     string    `json:"summary"`
	Number      int       `json:"number"`
	Step        string    `json:"step"`
	HealthScore float64   `json:"healthsore"`
	DishTypes   string    `json:"dishtypes"`
	Diets       string    `json:"diets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain food.Domain) Food {
	return Food {
		ID: domain.ID,
		Photo: domain.Photo,
		Summary: domain.Summary,
		Number: domain.Number,
		Step: domain.Step,
		HealthScore: domain.HealthScore,
		DishTypes:   domain.DishTypes,
		Diets: domain.Diets,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.CreatedAt,
	}
}

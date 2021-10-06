package response

import (
	"miniproject/business/food"
	"miniproject/business/recipesAPI"
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

type FoodAPI struct {
	ID          int       `json:"id"`
	Photo       string    `json:"photo"`
	Summary     string    `json:"summary"`
	Number      int       `json:"number"`
	Step        string    `json:"step"`
	HealthScore float64   `json:"healthsore"`
	DishTypes   string    `json:"dishtypes"`
	Diets       string    `json:"diets"`
}

func FromDomainAPI(domain recipesAPI.Domain) FoodAPI {
	return FoodAPI {
		ID: domain.ID,
		Photo: domain.Photo,
		Summary: domain.Summary,
		Number: domain.Number,
		Step: domain.Step,
		HealthScore: domain.HealthScore,
		DishTypes:   domain.DishTypes,
		Diets: domain.Diets,
	}
}
// func FromListDomain(record []food.Domain) (result []Food) {
// 	result = []Food{}
// 	for _, record := range record {
// 		result = append(result, FromDomain(record))
// 	}
// 	return result
// }
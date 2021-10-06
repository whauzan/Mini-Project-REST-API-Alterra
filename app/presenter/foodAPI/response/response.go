package response

import "miniproject/business/recipesAPI"

type FoodAPI struct {
	ID          int     `json:"id"`
	Photo       string  `json:"photo"`
	Summary     string  `json:"summary"`
	Number      int     `json:"number"`
	Step        string  `json:"step"`
	HealthScore float64 `json:"healthsore"`
	DishTypes   string  `json:"dishtypes"`
	Diets       string  `json:"diets"`
}

func FromDomainAPI(domain recipesAPI.Domain) FoodAPI {
	return FoodAPI{
		ID:          domain.ID,
		Photo:       domain.Photo,
		Summary:     domain.Summary,
		Number:      domain.Number,
		Step:        domain.Step,
		HealthScore: domain.HealthScore,
		DishTypes:   domain.DishTypes,
		Diets:       domain.Diets,
	}
}

package request

import "miniproject/business/food"

type Food struct {
	Name        string  `json:"name"`
	Photo       string  `json:"photo"`
	Summary     string  `json:"summary"`
	Number      int     `json:"number"`
	Step        string  `json:"step"`
	HealthScore float64 `json:"healthsore"`
	DishTypes   string  `json:"dishtypes"`
	Diets       string  `json:"diets"`
}

func ToDomain(request Food) food.Domain {
	return food.Domain{
		Name:        request.Name,
		Photo:       request.Photo,
		Summary:     request.Summary,
		Number:      request.Number,
		Step:        request.Step,
		HealthScore: request.HealthScore,
		DishTypes:   request.DishTypes,
		Diets:       request.Diets,
	}
}

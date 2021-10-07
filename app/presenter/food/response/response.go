package response

import (
	"miniproject/business/food"
	"time"
)

// type Food struct {
// 	ID          int       `json:"id"`
// 	Photo       string    `json:"photo"`
// 	Summary     string    `json:"summary"`
// 	Number      int       `json:"number"`
// 	Step        string    `json:"step"`
// 	HealthScore float64   `json:"healthsore"`
// 	DishTypes   string    `json:"dishtypes"`
// 	Diets       string    `json:"diets"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

// func FromDomain(domain food.Domain) Food {
// 	return Food {
// 		ID: domain.ID,
// 		Photo: domain.Photo,
// 		Summary: domain.Summary,
// 		Number: domain.Number,
// 		Step: domain.Step,
// 		HealthScore: domain.HealthScore,
// 		DishTypes:   domain.DishTypes,
// 		Diets: domain.Diets,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.CreatedAt,
// 	}
// }

// func FromListDomain(record []food.Domain) (result []Food) {
// 	result = []Food{}
// 	for _, record := range record {
// 		result = append(result, FromDomain(record))
// 	}
// 	return result
// }

type FoodInsertResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id"`
	Photo       string    `json:"photo"`
	Summary     string    `json:"summary"`
	HealthScore float64   `json:"healthscore"`
	DishTypes   string    `json:"dishtypes"`
	Diets       string    `json:"diets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainInsert(domain food.Domain) FoodInsertResponse {
	return FoodInsertResponse{
		Message:     "Insert Media Succes",
		ID:          domain.ID,
		Photo:       domain.Photo,
		Summary:     domain.Summary,
		HealthScore: domain.HealthScore,
		DishTypes:   domain.DishTypes,
		Diets:       domain.Diets,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

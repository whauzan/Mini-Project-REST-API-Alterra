package response

import "miniproject/business/foods"

type GetFoodByNameResponse struct {
	Result []struct {
		ID                   int      `json:"id"`
		Title                string   `json:"title"`
		Image                string   `json:"image"`
		Summary              string   `json:"summary"`
		HealthScore          float64  `json:"healthScore"`
		DishTypes            []string `json:"dishTypes"`
		Diets                []string `json:"diets"`
		AnalyzedInstructions []struct {
			Steps []struct {
				Number int    `json:"number"`
				Step   string `json:"step"`
			} `json:"steps"`
		} `json:"analyzedInstructions"`
	} `json:"results"`
}

func FromDomainFood(domain foods.Domain) GetFoodByNameResponse {
	return GetFoodByNameResponse{
		ID: domain.ID,
		
	}
}
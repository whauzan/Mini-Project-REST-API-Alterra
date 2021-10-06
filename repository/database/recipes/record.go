package recipes

import (
	"miniproject/business/recipesAPI"
)

type RecipeSource struct {
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

type Food struct {
	ID          int
	Title       string
	Image       string
	Summary     string
	DishTypes   string
	Diets       string
	Number      int
	Step        string
	HealthScore float64
}

func toDomain(record *Food) recipesAPI.Domain {
	return recipesAPI.Domain{
		ID:          record.ID,
		Name:        record.Title,
		Photo:       record.Image,
		Summary:     record.Summary,
		DishTypes:   record.DishTypes,
		Diets:       record.Diets,
		Number:      record.Number,
		Step:        record.Step,
		HealthScore: record.HealthScore,
	}
}

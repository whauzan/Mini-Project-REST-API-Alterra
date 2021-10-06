package recipes

import (
	// "encoding/json"
	"miniproject/business/recipesAPI"
	"strings"
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

type FoodAPI struct {
	ID          int
	Name        string
	Photo       string
	Summary     string
	Number      int
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
}

func toListDomain(record RecipeSource) []recipesAPI.Domain {
	var tempFood []recipesAPI.Domain
	var step string
	for _, value := range record.Result {
		for _, stepValue := range value.AnalyzedInstructions[0].Steps {
			step = string(string(rune(stepValue.Number))+". "+stepValue.Step+" ")
		}
		// steps, _ := json.Marshal(value.AnalyzedInstructions[0].Steps)
		tempFood = append(tempFood, recipesAPI.Domain {
		Name: value.Title, 
		Photo: value.Image, 
		Summary: value.Summary,
		DishTypes: strings.Join(value.DishTypes[:],", "),
		Diets: strings.Join(value.Diets[:], ", "),
		// Step: string(steps),
		HealthScore: value.HealthScore,
		Step: step})
	}
	return tempFood
}
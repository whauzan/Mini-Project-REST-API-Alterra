package recipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"miniproject/business/recipesAPI"
	"net/http"
)

type spoonacularAPI struct {
	client http.Client
}

func NewFoodAPI() recipesAPI.Repository {
	return &spoonacularAPI{
		client: http.Client{},
	}
}

func (s spoonacularAPI) GetRecipeByName(name string) (*recipesAPI.Domain, error) {
	apiKey := "dd6eecd6e5694ba2af2e94916aeed314"
	endpoint := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&addRecipeInformation=True&", apiKey, name)
	log.Println(endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	food := RecipeSource{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}

	var tempFood []struct {
		Title       string
		Image       string
		Summary     string
		DishTypes   string
		Diets       string
		Number      int
		Step        string
		HealthScore float64
	}
	for _, value := range food.Result {
		for _, valueStep := range value.AnalyzedInstructions[0].Steps {
			tempFood = append(tempFood, struct {
				Title       string
				Image       string
				Summary     string
				DishTypes   string
				Diets       string
				Number      int
				Step        string
				HealthScore float64
			}{Title: value.Title, Image: value.Image, Summary: value.Summary, Number: valueStep.Number, Step: valueStep.Step, HealthScore: value.HealthScore})
		}
		for _, valueDishType := range value.DishTypes {
			tempFood = append(tempFood, struct {
				Title       string
				Image       string
				Summary     string
				DishTypes   string
				Diets       string
				Number      int
				Step        string
				HealthScore float64
			}{DishTypes: valueDishType})
		}
		for _, valDiets := range value.Diets {
			tempFood = append(tempFood, struct {
				Title       string
				Image       string
				Summary     string
				DishTypes   string
				Diets       string
				Number      int
				Step        string
				HealthScore float64
			}{Diets: valDiets})
		}
	}
	// var titleFood string
	// var imageFood string
	// var summaryFood string
	// var dishType string
	// var dietFood string
	// var number int
	// var step string
	// var healthScore float64
	// var valueFood []
	// for _, v := range tempFood {
	// 	valueFood = append(valueFood, v.Title, v.Image, v.Summary, v.DishTypes, v.Diets, v.Number, v.Step, v.HealthScore)
	// }
	// titleFood = valueFood[0]
	// imageFood = valueFood[1]
	// step = valueFood[2]

	// res := toDomain(&Foods{Title: titleFood, Image: imageFood, Step: step})

	// return &res, nil
}

// func (s spoonacularAPI) GetHealthyRecipe(healthy string) (*foods.Domain, error) {

// }

// func (s spoonacularAPI) GetRecipeByCuisine(cuisine string) (*foods.Domain, error) {

// }

// func (s spoonacularAPI) GetRecipeByDiet(diet string) (*foods.Domain, error) {

// }

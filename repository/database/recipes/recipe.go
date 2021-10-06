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
	return spoonacularAPI{
		client: http.Client{},
	}
}

func (s spoonacularAPI) GetRecipeAPI(name string) ([]*recipesAPI.Domain, error) {
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

	return toListDomain(&food), nil
}
// func (s spoonacularAPI) GetHealthyRecipe(healthy string) (*foods.Domain, error) {

// }

// func (s spoonacularAPI) GetRecipeByCuisine(cuisine string) (*foods.Domain, error) {

// }

// func (s spoonacularAPI) GetRecipeByDiet(diet string) (*foods.Domain, error) {

// }
package food

import (
	"miniproject/business/recipesAPI"
)

type serviceFood struct {
	repository    Repository
	recipeAPIRepo recipesAPI.Repository
}

func NewService(repositoryFood Repository, recipeAPIRepo recipesAPI.Repository) Service {
	return &serviceFood{
		repository:    repositoryFood,
		recipeAPIRepo: recipeAPIRepo,
	}
}

func (service *serviceFood) GetRecipeByName(name string) (*Domain, error) {
	result, err := service.repository.GetRecipeByName(name)
	if err != nil {
		apiRecipe, err := service.GetRecipeAPI(name)
		if err != nil {
			return &Domain{}, err
		}
		insert, err := service.SaveFood(apiRecipe)
		if err != nil {
			return &Domain{}, nil
		}
		return insert, nil
	}
	return result, nil
}

func (service *serviceFood) GetRecipeAPI(name string) (*Domain, error) {
	result, err := service.recipeAPIRepo.GetRecipeByName(name)
	if err != nil {
		return &Domain{}, nil
	}
	newResult := Domain{
		ID:          result.ID,
		Name:        result.Name,
		Photo:       result.Photo,
		Summary:     result.Summary,
		DishTypes:   result.DishTypes,
		Diets:       result.Diets,
		Number:      result.Number,
		Step:        result.Step,
		HealthScore: result.HealthScore,
	}
	return &newResult, nil
}

func (s *serviceFood) SaveFood(food *Domain) (*Domain, error) {
	result, err := s.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

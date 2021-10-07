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


// func (service *serviceFood) GetAll() ([]Domain, error) {

// 	media, _ := service.repository.GetAll()
// 	if media == nil {
// 		return nil, business.ErrNotFound
// 	}

// 	return media, nil
// }

// func (service *serviceFood) GetByID(id int) (Domain, error) {

// 	media, err := service.repository.GetByID(id)
// 	if err != nil {
// 		return Domain{}, business.ErrNotFound
// 	}

// 	return media, nil
// }

func (service *serviceFood) SaveFood(name string, id int) (Domain, error) {
	APIFood, err := service.recipeAPIRepo.GetRecipeAPI(name)
	var result Domain
	if err != nil {
		for _, value := range APIFood {
			newRes := Domain{
				ID:          value.ID,
				Name:        value.Name,
				Photo:       value.Photo,
				Summary:     value.Summary,
				Step:        value.Step,
				HealthScore: value.HealthScore,
				DishTypes:   value.DishTypes,
				Diets:       value.Diets,
			}
			if newRes.ID == id {
				result, err = service.repository.Insert(&newRes)
				if err != nil {
					return Domain{}, err
				}
				return result, nil
			}
		}
	}
	return result, nil
}

package food

import (
	"miniproject/business/recipesAPI"
	"miniproject/business"
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

// type MediaService struct {
// 	repository     Repository
// 	contextTimeout time.Duration
// }

// func NewMediaService(repo Repository, timeout time.Duration) Service {
// 	return &MediaService{
// 		repository:     repo,
// 		contextTimeout: timeout,
// 	}
// }

// Business logic for medias crud
func (service *serviceFood) Create(domain *Domain) (Domain, error) {

	media, err := service.repository.Create(domain)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return media, nil
}

func (service *serviceFood) GetAll() ([]Domain, error) {

	media, _ := service.repository.GetAll()
	if media == nil {
		return nil, business.ErrNotFound
	}

	return media, nil
}

func (service *serviceFood) GetByID(id int) (Domain, error) {

	media, err := service.repository.GetByID(id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return media, nil
}

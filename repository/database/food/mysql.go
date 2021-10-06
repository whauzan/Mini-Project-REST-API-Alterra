package food

import (
	"miniproject/business/food"
	"miniproject/business/recipesAPI"

	"gorm.io/gorm"
)

type repositoryFood struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) food.Repository {
	return &repositoryFood{
		DB: db,
	}
}

func (repository *repositoryFood) GetRecipeByName(name string) ([]food.Domain, error) {
	// recordFood := Food{}
	var food []Food
	res := repository.DB.Where("name LIKE ?", "%"+name+"%")
	if res.Error != nil {
		return nil, res.Error
	}
	return toListDomain(food), nil
	// if err := repository.DB.Where("name LIKE ?", "%"+name+"%").Error; err != nil {
	// 	return []*food.Domain{}, err
	// }
	// var result  = toListDomain(recordFood)
	// return result, nil
}

func (repository *repositoryFood) Insert(foodie recipesAPI.Domain) (food.Domain, error) {
	// recordFood := fromDomain(*food)
	var food Food
	res := repository.DB.Create(&foodie)
	if res.Error != nil {
		return toDomain(food), res.Error
	}
	// if err := repository.DB.Create(&recordFood).Error; err != nil {
	// 	return food.Domain{}, err
	// }
	// result := toDomain(recordFood)
	// return &result, nil
	return toDomain(food), nil
}

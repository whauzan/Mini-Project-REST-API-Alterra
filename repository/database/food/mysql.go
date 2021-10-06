package food

import (
	"miniproject/business/food"

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

func (repository repositoryFood) GetRecipeByName(name string) (*food.Domain, error) {
	recordFood := Food{}
	if err := repository.DB.WHERE("name LIKE ?", "%"+name+"%").First(&recordFood).Error; err != nil {
		return &food.Domain{}, err
	}
	result := toDomain(recordFood)
	return result, nil
}

func (repository repositoryFood) Insert(food *food.Domain) (*food.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Create(&recordFood).Error; err != nil {
		return &food.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}
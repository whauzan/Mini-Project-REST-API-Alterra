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

func (rep *repositoryFood) Create(domain *food.Domain) (food.Domain, error) {
	foodie := fromDomain(*domain)
	if foodie.ID == 0 {
		return food.Domain{}, nil
	}
	result := rep.DB.Create(&foodie)

	if result.Error != nil {
		return food.Domain{}, result.Error
	}

	return toDomain(foodie), nil
}

func (rep *repositoryFood) GetAll() ([]food.Domain, error) {
	var foodie []Food
	result := rep.DB.Find(&foodie)

	if result.Error != nil {
		return nil, result.Error
	}

	return toListDomain(foodie), nil
}

func (rep *repositoryFood) GetByID(id int) (food.Domain, error) {
	var foodie Food
	result := rep.DB.First(&foodie, "ID = ?", id)

	if result.Error != nil {
		return food.Domain{}, result.Error
	}

	return toDomain(foodie), nil
}

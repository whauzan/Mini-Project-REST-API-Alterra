package food

import (
	"time"
)

type Domain struct {
	ID          int
	Name        string
	Photo       string
	Summary     string
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	SaveFood(name string, id int) (Domain, error)
}

type Repository interface {
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	Insert(food *Domain) (Domain, error)
}

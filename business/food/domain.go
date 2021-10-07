package food

import (
	"time"
)

type Domain struct {
	ID          int
	Name        string
	Photo       string
	Summary     string
	Number      int
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	Create(domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id int) (Domain, error)
}

type Repository interface {
	Create(domain *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id int) (Domain, error)
}

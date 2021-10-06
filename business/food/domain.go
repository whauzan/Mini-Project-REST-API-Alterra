package food

import "time"

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
	GetRecipeByName(name string) (*Domain, error)
	GetRecipeAPI(name string) (*Domain, error)
	SaveFood(food *Domain) (*Domain, error)
	// GetHealthyRecipe(limit int) (*Domain, error)
}

type Repository interface {
	GetRecipeByName(name string) (*Domain, error)
	Insert(food *Domain) (*Domain, error)
}

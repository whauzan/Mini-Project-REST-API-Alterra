package foods

type Domain struct {
	Name string
	Photo string
	Recipe []struct {
		Steps []struct {
			Number int
			Step string
		}
	}
	HealthScore float64
}

type Repository interface {
	GetRecipeByName(name string) (*Domain, error)
	GetHealthyRecipe() (*Domain, error)
	GetRecipeByCuisine(cuisine string) (*Domain, error)
	GetRecipeByDiet(diet string) (*Domain, error)
}
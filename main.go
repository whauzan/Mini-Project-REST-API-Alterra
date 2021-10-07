package main

import (
	"log"
	_foodHandler "miniproject/app/presenter/food"
	_foodAPIHandler "miniproject/app/presenter/foodAPI"
	"miniproject/app/routes"
	_foodService "miniproject/business/food"
	_foodRepo "miniproject/repository/database/food"
	_apiRepo "miniproject/repository/database/recipes"
	_dbDriver "miniproject/repository/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_foodRepo.Food{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: "root",
		DB_Password: "whr1728",
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Database: "miniproject",
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	e := echo.New()

	apiRepo := _apiRepo.NewFoodAPI()
	foodAPIHandler := _foodAPIHandler.NewFoodAPIHandler(apiRepo)

	foodRepo := _foodRepo.NewRepositoryMySQL(db)
	foodService := _foodService.NewService(foodRepo, apiRepo)
	foodHandler := _foodHandler.NewFoodHandler(foodService)

	routesInit := routes.HandlerList{
		FoodAPIHandler: *foodAPIHandler,
		FoodHandler:    *foodHandler,
	}
	routesInit.RouteReg(e)
	log.Fatal(e.Start(":8000"))
}

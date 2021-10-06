package main

import (
	"log"
	"miniproject/app/routes"
	_foodRepo "miniproject/repository/database/food"
	_apiRepo "miniproject/repository/database/recipes"
	_dbDriver "miniproject/repository/mysql"
	_foodService "miniproject/business/food"
	_foodHandler "miniproject/app/presenter/food"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_apiRepo.Food{},
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
	db := configDB.InitalDB()
	dbMigrate(db)

	e := echo.New()

	apiRepo := _apiRepo.NewFoodAPI()
	foodRepo := _foodRepo.NewRepositoryMySQL(db)
	foodService := _foodService.NewService(foodRepo, apiRepo)
	foodHandler := _foodHandler.NewHandler(foodService)

	routesInit := routes.HandlerList{
		FoodHandler: *foodHandler,
	}
	routesInit.RouteReg(e)
	log.Fatal(e.Start(":8000"))
}

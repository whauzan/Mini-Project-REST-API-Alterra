package main

import (
	"log"
	_foodAPIHandler "miniproject/app/presenter/foodAPI"
	"miniproject/app/routes"
	_apiRepo "miniproject/repository/database/recipes"

	"github.com/labstack/echo/v4"
)

// func dbMigrate(db *gorm.DB) {
// 	db.AutoMigrate(
// 		&_foodRepo.Food{},
// 	)
// }

func main() {
	// configDB := _dbDriver.ConfigDB{
	// 	DB_Username: "root",
	// 	DB_Password: "whr1728",
	// 	DB_Host:     "localhost",
	// 	DB_Port:     "3306",
	// 	DB_Database: "miniproject",
	// }
	// db := configDB.InitialDB()
	// dbMigrate(db)

	e := echo.New()

	apiRepo := _apiRepo.NewFoodAPI()
	// foodRepo := _foodRepo.NewRepositoryMySQL(db)
	// foodService := _foodService.NewService(foodRepo, apiRepo)
	foodAPIHandler := _foodAPIHandler.NewFoodAPIHandler(apiRepo)

	routesInit := routes.HandlerList{
		FoodAPIHandler: *foodAPIHandler,
	}
	routesInit.RouteReg(e)
	log.Fatal(e.Start(":8000"))
}

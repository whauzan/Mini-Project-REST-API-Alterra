package main

import (
	"log"

	_userService "miniproject/business/users"
	_userHandler "miniproject/app/presenter/user"
	_driverFactory "miniproject/repository"
	_userRepo "miniproject/repository/database/users"
	_foodAPIHandler "miniproject/app/presenter/foodAPI"
	_apiRepo "miniproject/repository/database/recipes"
	_dbDriver "miniproject/repository/mysql"
	_foodService "miniproject/business/food"
	_foodHandler "miniproject/app/presenter/food"
	_foodRepo "miniproject/repository/database/food"
	_middleware "miniproject/app/middleware"
	_routes "miniproject/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_foodRepo.Food{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB {
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host: viper.GetString(`database.host`),
		DB_Port: viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       "rahasiaayam",
		ExpiresDuration: 1,
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, 10, &configJWT)
	userHandler := _userHandler.NewUserHandler(userService)
	apiRepo := _apiRepo.NewFoodAPI()
	foodAPIHandler := _foodAPIHandler.NewFoodAPIHandler(apiRepo)
	foodRepo := _foodRepo.NewRepositoryMySQL(db)
	foodService := _foodService.NewService(foodRepo, apiRepo)
	foodHandler := _foodHandler.NewFoodHandler(foodService)

	routesInit := _routes.HandlerList{
		UserHandler: *userHandler,
		FoodAPIHandler: *foodAPIHandler,
		FoodHandler:    *foodHandler,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}

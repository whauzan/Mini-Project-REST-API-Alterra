package main

import (
	"log"

	_userService "miniproject/business/users"
	_userController "miniproject/controller/users"
	_driverFactory "miniproject/repository"
	_userRepo "miniproject/repository/database/users"

	_dbDriver "miniproject/repository/mysql"

	_middleware "miniproject/app/middleware"
	_routes "miniproject/app/routes"

	"github.com/labstack/echo/v4"
	// "github.com/spf13/viper"
	"gorm.io/gorm"
)

// func init() {
// 	viper.SetConfigFile(`app/config/config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}

// 	if viper.GetBool(`debug`) {
// 		log.Println("Service RUN on DEBUG mode")
// 	}
// }

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
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

	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       "rahasiaayam",
		ExpiresDuration: 1,
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, 10, &configJWT)
	userCtrl := _userController.NewUserController(userService)

	routesInit := _routes.ControllerList{
		UserController: *userCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}

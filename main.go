package main

import (
	_apiRepo "miniproject/repository/database/recipe"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_apiRepo.Foods{},
	)
}

func main() {
	
}

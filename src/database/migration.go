package database

import (
	"digiexam-interview/pkg/postgres"
	"digiexam-interview/src/models"
	"fmt"
)

func MigrationUp() {
	err := postgres.DB.AutoMigrate(
		&models.GradingScale{},
		&models.GradingScaleMeta{},
		&models.School{},
		&models.Course{},
		&models.Student{},
		&models.Result{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migrations Updated.")
}
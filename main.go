package main

import (
	"digiexam-interview/pkg/postgres"
	"digiexam-interview/src/database"
	"digiexam-interview/src/database/seeder"
	"digiexam-interview/src/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Connect to Database
	postgres.InitialiseDatabaseConnection()

	// Run migration
	database.MigrationUp()

	// Seed Data to Database
	seeder.Seed(*postgres.DB,seeder.SeedDatabaseInput{
		SchoolCount: 2,
		CourseCount: 10,
		StudentCount: 10,
	})
	
	router := mux.NewRouter()
	routes.RouteInit(router.PathPrefix("/api").Subrouter())

	port := os.Getenv("SERVER_PORT");
	host := os.Getenv("SERVER_HOST");

	fmt.Printf("Server Running on %s:%s\n",host,port)

	http.ListenAndServe(fmt.Sprintf("%s:%s",host,port), router)
}

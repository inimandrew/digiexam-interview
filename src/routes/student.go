package routes

import (
	"digiexam-interview/pkg/postgres"
	"digiexam-interview/src/controllers"
	"digiexam-interview/src/repositories"

	"github.com/gorilla/mux"
)

func StudentRoutes(r *mux.Router){
	studentRepository := repositories.RepositoryStudent(postgres.DB)
	c := controllers.StudentHandler(studentRepository);

	r.HandleFunc("/students/gpa", c.FetchStudents).Methods("GET")
}
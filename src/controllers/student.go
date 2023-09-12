package controllers

import (
	"digiexam-interview/src/repositories"
	"digiexam-interview/src/responses"
	"encoding/json"
	"net/http"
	"strconv"
)

type StudentController struct {
	StudentRepository repositories.StudentRepository
}

func StudentHandler(studentRepository repositories.StudentRepository) *StudentController {
	return &StudentController{studentRepository}
}

func (sc *StudentController) FetchStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var page int = 1
	var limit int = 10

	if len(r.URL.Query().Get("page")) != 0 {
		pageValue, _ := strconv.Atoi(r.URL.Query().Get("page"))
		page = pageValue;
	}

	if len(r.URL.Query().Get("limit")) != 0 {
		limitValue, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		limit = limitValue
	}
	
	students,meta, err := sc.StudentRepository.FetchStudents(page,limit)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := responses.PaginatedResults{Code: http.StatusOK, Data: students, PaginateMeta: meta}
	json.NewEncoder(w).Encode(response)
}
package repositories

import (
	"digiexam-interview/src/models"
	"digiexam-interview/src/responses"
	"digiexam-interview/src/services"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchStudents(page int, limit int) ([]responses.Student,responses.PaginateMeta, error)
}

func RepositoryStudent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FetchStudents(page int, limit int) ([]responses.Student, responses.PaginateMeta, error) {
	var students []models.Student
	var total int64
	offset := (page - 1) * limit;

	err := r.db.Limit(limit).Offset(offset).Preload("School.GradingScale.GradingScaleMetas").Preload("Results.Course").Find(&students).Error
	r.db.Model(&models.Student{}).Count(&total);

	meta := services.GetPaginationMeta(total,int64(len(students)),int64(limit),int64(offset))
	formattedStudents  := make([]responses.Student, 0,len(students));

	for _, student := range students {
		formattedStudents = append(formattedStudents, student.Format())
	}
	
	return formattedStudents, meta, err
}
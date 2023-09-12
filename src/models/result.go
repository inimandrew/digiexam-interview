package models

import (
	"digiexam-interview/src/responses"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Result struct {
	BaseModel
	StudentID uuid.UUID		`gorm:"foreignKey: ID;not null"`
	Student Student
	CourseID uuid.UUID		`gorm:"foreignKey: ID;not null"`
	Course Course
	Percentage int		`gorm:"type: int;not null"`
	Grade string		`gorm:"type: varchar;size:3;not null"`
}

func (r Result) Format() responses.Result {
	result :=  responses.Result{
		ID: r.ID.String(),
		CreatedAt: r.CreatedAt.Format(time.RFC822),
		UpdatedAt: r.UpdatedAt.Format(time.RFC822),
		Percentage: r.Percentage,
		Course: r.Course.Format(),
		Grade: r.Grade,
	}

	return result
}
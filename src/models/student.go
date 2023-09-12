package models

import (
	"digiexam-interview/src/responses"
	"digiexam-interview/src/services"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Student struct {
	BaseModel
	FirstName string	`gorm:"type: varchar;size: 255;not null"`
	LastName string		`gorm:"type: varchar;size: 255;not null"`
	Email string 		`gorm:"type: varchar;size: 255;unique;not null"`
	SchoolID uuid.UUID	`gorm:"foreignKey: ID;not null"`
	School School
	Results []Result	`gorm:"foreignKey: StudentID"`
}

func (s Student) Format() responses.Student {
	results := make([]responses.Result,0);
	gradingScaleMetas := make([]responses.GradingScaleMeta,0);

	for _, result := range s.Results{
		results = append(results, result.Format())
	}

		for _, gradingScaleMeta := range s.School.GradingScale.GradingScaleMetas{
		gradingScaleMetas = append(gradingScaleMetas, gradingScaleMeta.Format())
	}

	student :=  responses.Student{
		ID: s.ID.String(),
		CreatedAt: s.CreatedAt.Format(time.RFC822),
		UpdatedAt: s.UpdatedAt.Format(time.RFC822),
		FullName: fmt.Sprintf("%s %s",s.FirstName, s.LastName),
		Email: s.Email,
		Results: results,
		School: s.School.Format(),
		Gpa: services.CalculateGpa(results,gradingScaleMetas),
	}

	return student
}
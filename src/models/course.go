package models

import (
	"digiexam-interview/src/responses"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Course struct {
	BaseModel
	Name string				`gorm:"type: varchar;size: 255;not null"`
	SchoolID uuid.UUID		`gorm:"foreignKey: ID;not null"`
	Credit int				`gorm:"type: int;default: 1;not null"`
	Results []Result		`gorm:"foreignKey: CourseID"`
}

func (c Course) Format() responses.Course {
	course :=  responses.Course{
		ID: c.ID.String(),
		CreatedAt: c.CreatedAt.Format(time.RFC822),
		UpdatedAt: c.UpdatedAt.Format(time.RFC822),
		Name: c.Name,
		Credit: c.Credit,
	}

	return course
}
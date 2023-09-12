package models

import (
	"digiexam-interview/src/responses"
	"time"

	uuid "github.com/satori/go.uuid"
)

type School struct {
	BaseModel
	Name string					`gorm:"type: varchar;size: 255;not null"`
	GradingScaleID uuid.UUID			`gorm:"foreignKey: ID;not null"`
	GradingScale GradingScale
	Courses []Course			`gorm:"foreignKey: SchoolID"`
	Students []Student			`gorm:"foreignKey: SchoolID"`
}

func (s School) Format() responses.School {
	school :=  responses.School{
		ID: s.ID.String(),
		CreatedAt: s.CreatedAt.Format(time.RFC822),
		UpdatedAt: s.UpdatedAt.Format(time.RFC822),
		Name: s.Name,
		GradingScale: s.GradingScale.Format(),
	}

	return school
}
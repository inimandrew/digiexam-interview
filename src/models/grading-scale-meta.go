package models

import (
	"digiexam-interview/src/responses"
	"time"

	uuid "github.com/satori/go.uuid"
)

type GradingScaleMeta struct {
	BaseModel
	GradingScaleID uuid.UUID	`gorm:"foreignKey: ID"`
	MinimumPercentage int		`gorm:"type: int;not null"`
	MaximumPercentage int		`gorm:"type: int;not null"`
	Grade string				`gorm:"type: char;size: 2;not null"`
	GpaValue int				`gorm:"type: int;not null"`
}

func (gsm GradingScaleMeta) Format() responses.GradingScaleMeta {
	gradingScaleMeta :=  responses.GradingScaleMeta{
		ID: gsm.ID.String(),
		CreatedAt: gsm.CreatedAt.Format(time.RFC822),
		UpdatedAt: gsm.UpdatedAt.Format(time.RFC822),
		MinimumPercentage: gsm.MinimumPercentage,
		MaximumPercentage: gsm.MaximumPercentage,
		Grade: gsm.Grade,
		GpaValue: gsm.GpaValue,
	}

	return gradingScaleMeta
}
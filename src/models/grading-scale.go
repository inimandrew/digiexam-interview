package models

import (
	"digiexam-interview/src/responses"
	"time"
)

type GradingScale struct {
	BaseModel
	Name string								`gorm:"type: varchar;size: 255;unique:;not null"`
	GradingScaleMetas []GradingScaleMeta	`gorm:"foreignKey: GradingScaleID"`
	Schools []School						`gorm:"foreignKey: GradingScaleID"`
}

func (g GradingScale) Format() responses.GradingScale {
	gradingScale :=  responses.GradingScale{
		ID: g.ID.String(),
		CreatedAt: g.CreatedAt.Format(time.RFC822),
		UpdatedAt: g.UpdatedAt.Format(time.RFC822),
		Name: g.Name,
	}

	return gradingScale
}
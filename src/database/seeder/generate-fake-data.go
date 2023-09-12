package seeder

import "digiexam-interview/src/models"

func GetFakeGradingScale() []models.GradingScale {
	gradingScales := []models.GradingScale{
		{
			Name: "America Grading Scale",
			GradingScaleMetas: []models.GradingScaleMeta{
				{
					MinimumPercentage: 91,
					MaximumPercentage: 100,
					Grade:             "A",
					GpaValue:          4,
				},
				{
					MinimumPercentage: 80,
					MaximumPercentage: 90,
					Grade:             "B",
					GpaValue:          3,
				},
				{
					MinimumPercentage: 70,
					MaximumPercentage: 79,
					Grade:             "C",
					GpaValue:          2,
				},
				{
					MinimumPercentage: 60,
					MaximumPercentage: 69,
					Grade:             "D",
					GpaValue:          1,
				},
				{
					MinimumPercentage: 0,
					MaximumPercentage: 59,
					Grade:             "F",
					GpaValue:          0,
				},
			},
		},
		{
			Name: "Nigeria Grading Scale",
			GradingScaleMetas: []models.GradingScaleMeta{
				{
					MinimumPercentage: 75,
					MaximumPercentage: 100,
					Grade:             "A",
					GpaValue:          5,
				},
				{
					MinimumPercentage: 70,
					MaximumPercentage: 74,
					Grade:             "B",
					GpaValue:          4,
				},
				{
					MinimumPercentage: 60,
					MaximumPercentage: 69,
					Grade:             "C",
					GpaValue:          3,
				},
				{
					MinimumPercentage: 50,
					MaximumPercentage: 59,
					Grade:             "D",
					GpaValue:          2,
				},
				{
					MinimumPercentage: 40,
					MaximumPercentage: 49,
					Grade:             "E",
					GpaValue:          1,
				},
				{
					MinimumPercentage: 39,
					MaximumPercentage: 0,
					Grade:             "F",
					GpaValue:          0,
				},
			},
		},
	}
	return gradingScales
}
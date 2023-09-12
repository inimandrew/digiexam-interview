package services_test

import (
	"digiexam-interview/src/database/seeder"
	"digiexam-interview/src/models"
	"digiexam-interview/src/responses"
	"digiexam-interview/src/services"
	"testing"

	"github.com/thoas/go-funk"
)

type CalculateGradeInput struct {
	percentage int
	gradingScaleMetas []responses.GradingScaleMeta
	expectedGrade string
}

type CalculateGpaInput struct {
	results []responses.Result
	gradingScaleMetas []responses.GradingScaleMeta
	gpa float64
}

func TestCalculcateGrade(t *testing.T){
	gradingScale := seeder.GetFakeGradingScale();
	americaGradingScaleMetas := funk.Map(gradingScale[0].GradingScaleMetas, func(u models.GradingScaleMeta) responses.GradingScaleMeta {
						return u.Format()
					}).([]responses.GradingScaleMeta);
	nigeriaGradingScaleMetas := funk.Map(gradingScale[1].GradingScaleMetas, func(u models.GradingScaleMeta) responses.GradingScaleMeta {
						return u.Format()
					}).([]responses.GradingScaleMeta);

	testData := []CalculateGradeInput{
		{
			91,
			americaGradingScaleMetas,
			"A",
		},
		{
			50,
			nigeriaGradingScaleMetas,
			"D",
		},
		{
			70,
			nigeriaGradingScaleMetas,
			"B",
		},
		{
			80,
			nigeriaGradingScaleMetas,
			"A",
		},
			{
			50,
			americaGradingScaleMetas,
			"F",
		},
	};

			for _, datum := range testData {
		grade := services.CalculateGrade(datum.percentage, datum.gradingScaleMetas);

			if grade == datum.expectedGrade {
		t.Logf("CalculateGrade(%d) passed successfully", datum.percentage,)
	}else{
		t.Errorf("CalculateGrade(%d) failed!", datum.percentage)
	}
	}
}

func TestCalculateGpa(t *testing.T){
	gradingScale := seeder.GetFakeGradingScale();
	americaGradingScaleMetas := funk.Map(gradingScale[0].GradingScaleMetas, func(u models.GradingScaleMeta) responses.GradingScaleMeta {
						return u.Format()
					}).([]responses.GradingScaleMeta);
	nigeriaGradingScaleMetas := funk.Map(gradingScale[1].GradingScaleMetas, func(u models.GradingScaleMeta) responses.GradingScaleMeta {
						return u.Format()
					}).([]responses.GradingScaleMeta);

		 nigerianStudentResults := []responses.Result{
			{
				Course: responses.Course{
					Credit: 4,
				},
				Percentage: 90,
				Grade: services.CalculateGrade(90, nigeriaGradingScaleMetas),
			},
			{
				Course: responses.Course{
					Credit: 5,
				},
				Percentage: 71,
				Grade: services.CalculateGrade(71, nigeriaGradingScaleMetas),
			},
			{
				Course: responses.Course{
					Credit: 2,
				},
				Percentage: 65,
				Grade: services.CalculateGrade(65, nigeriaGradingScaleMetas),
			},
		}

		americanStudentResults := []responses.Result{
			{
				Course: responses.Course{
					Credit: 4,
				},
				Percentage: 92,
				Grade: services.CalculateGrade(92, americaGradingScaleMetas),
			},
			{
				Course: responses.Course{
					Credit: 2,
				},
				Percentage: 85,
				Grade: services.CalculateGrade(85, americaGradingScaleMetas),
			},
			{
				Course: responses.Course{
					Credit: 3,
				},
				Percentage: 65,
				Grade: services.CalculateGrade(65, americaGradingScaleMetas),
			},
		}
	

		testData := []CalculateGpaInput{
			{
				nigerianStudentResults,
				nigeriaGradingScaleMetas,
				4.19,
			},
			{
				americanStudentResults,
				americaGradingScaleMetas,
				2.78,
			},
		};
			
		
		for _, datum := range testData {
					gpa := services.CalculateGpa(datum.results, datum.gradingScaleMetas);
			if gpa == datum.gpa {
		t.Logf("CalculateGpa passed successfully")
	}else{
		t.Errorf("CalculateGpa failed!")
	}
		}

}
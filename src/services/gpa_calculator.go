package services

import (
	"digiexam-interview/src/responses"
	"math"
)


func CalculateGpa(results []responses.Result, gradingScaleMetas []responses.GradingScaleMeta)float64 {	
	gradeGpaValueMap := make(map[string]int)
	var maximumCredit int
	var totalGpaProduct int

	for _, gradingScaleMeta := range gradingScaleMetas {
		    gradeGpaValueMap[gradingScaleMeta.Grade] = gradingScaleMeta.GpaValue
	}

	for _, result := range results {
		maximumCredit = maximumCredit + result.Course.Credit
		totalGpaProduct += result.Course.Credit * gradeGpaValueMap[result.Grade]
	}

	return math.Ceil(float64(totalGpaProduct) / float64(maximumCredit) * 100)/100
}

func CalculateGrade(percentage int, gradingScaleMetas []responses.GradingScaleMeta) string {
	var grade = "F"
	for _, gradingScaleMeta := range gradingScaleMetas {
		if percentage <= gradingScaleMeta.MaximumPercentage && percentage >= gradingScaleMeta.MinimumPercentage{
			grade = gradingScaleMeta.Grade
		}
	}

	return grade
}
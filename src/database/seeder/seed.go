package seeder

import (
	"digiexam-interview/src/models"
	"digiexam-interview/src/responses"
	"digiexam-interview/src/services"
	"fmt"
	"math/rand"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

type SeedDatabaseInput struct{
	SchoolCount int
	CourseCount int
	StudentCount int
}

func Seed(db gorm.DB, s SeedDatabaseInput) bool{
	gradingScales := GetFakeGradingScale()

	var persistedGradingScales []models.GradingScale
	 db.Find(&persistedGradingScales)

	if len(persistedGradingScales) == 0{
		db.Create(&gradingScales)
		db.Find(&persistedGradingScales)

		schools := make([]models.School,0, s.SchoolCount)
		schoolNames := make([]string,0)
		for i := 0; i < s.SchoolCount; i++ {
			randGradingScale := persistedGradingScales[rand.Intn(len(persistedGradingScales))];
			schoolName := faker.Company().Name();

			schools = append(schools, models.School{
				Name: schoolName,
				GradingScaleID: randGradingScale.ID,
			})

			schoolNames = append(schoolNames, schoolName)
		}
		
		db.Create(&schools)
	
		var persistedSchools []models.School
		db.Where("name IN ?", schoolNames).Find(&persistedSchools);
	
		students := make([]models.Student,0, s.StudentCount)
		courses := make([]models.Course, 0, s.CourseCount);
	
			for _,school := range persistedSchools {
			for i := 0; i < s.StudentCount; i++ {
				students = append(students, models.Student{
				FirstName: faker.Name().FirstName(),
				LastName: faker.Name().LastName(),
				Email: faker.Internet().Email(),
				SchoolID: school.ID,
			})
			}
	
			for i := 0; i < s.CourseCount; i++ {
				courses = append(courses, models.Course{
					Name: faker.Commerce().ProductName(),
					SchoolID: school.ID,
					Credit: faker.RandomInt(1,6),
				})
			}
		}
	
		db.Create(&students)
		db.Create(&courses)
	
		db.Preload("Courses").Preload("Students").Preload("GradingScale.GradingScaleMetas").Find(&persistedSchools);
		results := make([]models.Result, 0, len(students) * len(courses));
	
				for _,school := range persistedSchools {
					gradingScaleMetas :=  funk.Map(school.GradingScale.GradingScaleMetas, func(u models.GradingScaleMeta) responses.GradingScaleMeta {
						return u.Format()
					}).([]responses.GradingScaleMeta);

					schoolCourses := school.Courses;
					schoolStudents := school.Students;
	
					for _, student := range schoolStudents {
						for _, course := range schoolCourses {
							percentage :=faker.RandomInt(30,100);

						results = append(results, models.Result{
							CourseID: course.ID,
							StudentID: student.ID,
							Percentage: percentage,
							Grade: services.CalculateGrade(percentage,gradingScaleMetas),
						})
					}
					}
				}
	
			db.Create(&results)
			fmt.Println("Database Seeded Successfully")
	}

	return true;
}
package main

import (
	"github.com/ivanmrnn/uadmin_sample/student_information/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Student_Information{},
		models.School{},
		models.Course{},
	)
	uadmin.RegisterInlines(models.School{}, map[string]string{
		"Student_Information": "SchoolID",
	})
	uadmin.RegisterInlines(models.Course{}, map[string]string{
		"Student_Information": "CourseID",
	})
	uadmin.StartServer()
}

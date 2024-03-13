package main

import (
    "github.com/uadmin/uadmin"
	"github.com/ivanmrnn/uadmin_sample/student_information/models"
)

func main() {
	uadmin.Register(
        models.Student_Information{},
		models.Guardian{},
		models.School{},
		models.Course{},
    )
	uadmin.RegisterInlines(models.Guardian{}, map[string]string{
		"Student_Information": "GuardianID",
	})
	uadmin.RegisterInlines(models.School{}, map[string]string{
		"Student_Information": "SchoolID",
	})
	uadmin.RegisterInlines(models.Course{}, map[string]string{
		"Student_Information": "CourseID",
	})
    uadmin.StartServer()
}
package models

import (
	"fmt"
	"time"
	"github.com/uadmin/uadmin"
)


type Student_Information struct {
	uadmin.Model
	Student_Number uint
	Name           string `uadmin:"required;search;display_name:Student Name"`
	Student_ID     string `uadmin:"search;display_name:Student ID"`
	IdPicture      string `uadmin:"image"`
	DateOfBirth    time.Time
	Age            uint   `uadmin:"default_value: ;min:2;max:120"`
	Address        string `uadmin:"search"`
	Phone          string
	Email          string   `uadmin:"email"`
	FamilyIncome   uint `uadmin:"money"`
	School         School
	SchoolID       uint
	Course         Course `uadmin:"search"`
	CourseID       uint
}


func (s *Student_Information) Save() {
	currentYear := time.Now().Year() % 100
 	
	uadmin.Preload(s)
	code := s.Course.CourseCode
	s.Student_ID = generateStudentNumber(currentYear, code, s.Student_Number+1)
	uadmin.Save(s)
	//uadmin.Trail(uadmin.DEBUG,s.Test)
}

func generateStudentNumber(year int, code, studentID uint) string {
	return fmt.Sprintf("%d-%v%05d", year%100, code, studentID)
}

/*func generateBSU(year int, studentID uint) string {
	return fmt.Sprintf("%d-%05d", year, studentID)
}
func generateUB(year int, code uint, studentID uint) string {
	return fmt.Sprintf("%d%02d%04d", year, code, studentID)
}
func generateUP(code uint, year int) string {
	randomNum := rand.Intn(9999)
	return fmt.Sprintf("%02d%d%d", code, year, randomNum)
}*/

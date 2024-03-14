package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

type Student_Information struct {
	uadmin.Model
	CreatedAt      time.Time `uadmin:"hidden"`
	Student_Number uint
	Name           string `uadmin:"required;search;display_name:Student Name"`
	Student_ID     string `uadmin:"search;display_name:Student ID"`
	IdPicture      string `uadmin:"image"`
	DateOfBirth    time.Time
	Age            uint   `uadmin:"default_value: ;min:2;max:120"`
	Address        string `uadmin:"search"`
	Zip            uint   `uadmin:"default_value: ;search"`
	Phone          string
	Email          string   `uadmin:"email"`
	Guardian       Guardian `uadmin:"help:Parent or Guardian Name"`
	GuardianID     uint
	FamilyIncome   uint   `uadmin:"money"`
	School         School `uadmin:"search"`
	SchoolID       uint
	Course         Course `uadmin:"search"`
	CourseID       uint
}

func (s *Student_Information) Save() {
	studentNumber := generateStudentNumber(s.CreatedAt.Year(), s.Student_Number+1)
	s.Student_ID = studentNumber
	uadmin.Save(s)
}

func generateStudentNumber(year int, studentID uint) string {
	return fmt.Sprintf("%d-%05d", year%100, studentID)
}

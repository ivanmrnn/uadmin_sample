package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Student_Information struct {
	uadmin.Model
	Name           string `uadmin:"required;search;display_name:Student Name"`
	Student_number string `uadmin:"search;display_name:Student ID"`
	IdPicture      string `uadmin:"image"`
	DateOfBirth    time.Time
	Age            uint   `uadmin:"min:2;max:120"`
	Address        string `uadmin:"search"`
	Zip            uint   `uadmin:"search"`
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

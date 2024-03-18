package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/uadmin/uadmin"
)

type Student_Information struct {
	uadmin.Model
	Name         string `uadmin:"required;search;display_name:Student Name"`
	Student_ID   string `uadmin:"search;display_name:Student ID;read_only"`
	IdPicture    string `uadmin:"image"`
	DateOfBirth  time.Time 
	Age          uint   `uadmin:"default_value: ;min:2;max:120"`
	Address      string
	Phone        string
	Email        string `uadmin:"email"`
	FamilyIncome uint   `uadmin:"default_value: "`
	School       School `uadmin:"search;required"`
	SchoolID     uint
	SchoolName   string `uadmin:"hidden;list_exclude"`
	Course       Course `uadmin:"search;required"`
	CourseID     uint
	IDFormat     IDFormat `uadmin:"hidden;list_exclude"`
	IDFormatID   uint
}

func (s Student_Information) Validate() (errMsg map[string]string) {
	errMsg = map[string]string{}

	students := Student_Information{}
	if uadmin.Count(&students, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This student is already registered in the system."
	}
	return
}

func (s *Student_Information) Save() {
	uadmin.Preload(s)
	s.SchoolName = s.School.Name
	formatting := s.School.Format
	courseCode := s.Course.CourseCode
	schoolCode := s.School.SchoolCode
	currentYear := time.Now().Year()

	students := Student_Information{}
	studentID := s.Student_ID

	if studentID == "" {
		baseCount := uadmin.Count(&students, "school_name = ?", s.SchoolName)
		recentCount := baseCount + 1

		
		characters := []string{}
		generated := processSchoolFormat(characters , formatting, currentYear, courseCode, schoolCode, recentCount)
		
		s.Student_ID = generated
	} /*else {
		existingStudent := Student_Information{}
        if uadmin.Find(&existingStudent, "student_id = ?", studentID).Error == nil {
            s.Student_ID = existingStudent.Student_ID
        }
	}*/
	uadmin.Save(s)
}

func generateRandomNumber(year int) string {
	randomNum := rand.Intn(99999)
	return fmt.Sprintf("%v %v", year, randomNum)
}

func generateStudentNumber(count int, id int) string {
	generateFormat := fmt.Sprintf("%%0%dd", count)
	return fmt.Sprintf(generateFormat, id)
}

func processSchoolFormat(characters []string, formatting string, currentYear int, courseCode uint, schoolCode uint, recentCount int) string {
	generated := ""
	for i := 0; i < len(formatting); i++ {
		characters = append(characters, string(formatting[i]))
	}
	for len(characters) > 0 {
		switch characters[0] {
		case "Y":
			switch {
			case len(characters) >= 4 && characters[1] == "Y" && characters[2] == "Y" && characters[3] == "Y":
				generated += strconv.FormatUint(uint64(currentYear), 10)
				characters = characters[4:]
			case len(characters) >= 3 && characters[1] == "Y" && characters[2] == "Y":
				generated += strconv.FormatUint(uint64(currentYear%100), 10)
				characters = characters[3:]
			case len(characters) >= 2 && characters[1] == "Y":
				generated += strconv.FormatUint(uint64(currentYear%100), 10)
				characters = characters[2:]
			default:
				characters = characters[1:]
			}
		case "C":
			if len(characters) >= 2 && characters[1] == "C" {
				generated += strconv.FormatUint(uint64(courseCode), 10)
				characters = characters[2:]
			} else {
				characters = characters[1:]
			}
		case "N":
			count := 1
			for count < len(characters) && characters[count] == characters[0] {
				count++
			}
			generated += generateStudentNumber(count, recentCount)
			characters = characters[count:]
		case "R":
			count := 1
			repeat := ""
			for count < len(characters) && characters[count] == characters[0] {
				count++
			}
			nineCount := count
			for nineCount > 0 {
				repeat += "9"
				nineCount--
			}
			repeatFormat, _ := strconv.Atoi(repeat)
			randomNum := rand.Intn(repeatFormat)
			generated += strconv.FormatUint(uint64(randomNum), 10)
			characters = characters[count:]
		case "S":
			if len(characters) >= 2 && characters[1] == "S" {
				generated += strconv.FormatUint(uint64(schoolCode), 10)
				characters = characters[2:]
			} else {
				characters = characters[1:]
			}
		case " ":
			count := 1
			for count < len(characters) && characters[count] == characters[0] {
				count++
			}
			generated += strings.Repeat(characters[0], count)
			characters = characters[count:]
		default:
			year := currentYear % 100
			generated = generateRandomNumber(year)
			characters = []string{}
		}
	}
	return generated
}



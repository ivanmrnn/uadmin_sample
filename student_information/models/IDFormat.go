package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type IDFormat struct {
	uadmin.Model
	Name   string `uadmin:"hidden"`
	Format string `uadmin:"pattern:^[c,C,n,N,r,R,s,S,y,Y, ]*$;pattern_msg:Invalid Format;required;help:Format: YY or YYYY for year, CC for Course Code, SS for School Code, R for random number, N for student number. Must atleast have NNN"`
}

func (f *IDFormat) Save() {
	formatting := strings.ToUpper(f.Format)

	hasNNN := strings.Contains(formatting, "NNN")
	if !hasNNN {
		formatting = "YY NNNNN"
	}

	characters := []string{}
	formatted := processFormatCharacters(characters , formatting)	
	f.Format = formatted
	f.Name = f.Format
	uadmin.Save(f)
}

func processFormatCharacters(characters []string, formatting string) string {
	formatted := ""
	for i := 0; i < len(formatting); i++ {
		characters = append(characters, string(formatting[i]))
	}
	for len(characters) > 0 {
		switch characters[0] {
		case "Y":
			switch {
			case len(characters) >= 4 && characters[1] == "Y" && characters[2] == "Y" && characters[3] == "Y":
				formatted += "YYYY"
				characters = characters[4:]
			case len(characters) >= 3 && characters[1] == "Y" && characters[2] == "Y":
				formatted += "YY"
				characters = characters[3:]
			case len(characters) >= 2 && characters[1] == "Y":
				formatted += "YY"
				characters = characters[2:]
			default:
				characters = characters[1:]
			}
		case "C":
			if len(characters) >= 2 && characters[1] == "C" {
				formatted += "CC"
				characters = characters[2:]
			} else {
				characters = characters[1:]
			}
		case "S":
			if len(characters) >= 2 && characters[1] == "S" {
				formatted += "SS"
				characters = characters[2:]
			} else {
				characters = characters[1:]
			}
		case "R", "N", " ":
			count := 1
			for count < len(characters) && characters[count] == characters[0] {
				count++
			}
			formatted += strings.Repeat(characters[0], count)
			characters = characters[count:]
		default:
			characters = []string{}
		}
	}
	return formatted
}
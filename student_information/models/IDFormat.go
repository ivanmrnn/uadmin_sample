package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type IDFormat struct {
	uadmin.Model
	Name   string
	Format string `uadmin:"pattern:^[c,C,n,N,r,R,s,S,y,Y,-, ]*$;pattern_msg:Invalid Format"`
}

func (f *IDFormat) Save() {
	formatting := strings.ToUpper(f.Format)

	hasNNN := strings.Contains(formatting, "NNN")
	hasRRR := strings.Contains(formatting, "RRR")
	if !hasNNN && !hasRRR {
		formatting = "YY-RRRRR"
	}

	formatted := ""
	characters := []string{}
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
				formatted += "Y"
				characters = characters[1:]
			}
		case "C":
			if len(characters) >= 2 && characters[1] == "C" {
				formatted += "CC"
				characters = characters[2:]
			} else {
				formatted += "C"
				characters = characters[1:]
			}
		case "S":
			if len(characters) >= 2 && characters[1] == "S" {
				formatted += "SS"
				characters = characters[2:]
			} else {
				formatted += "S"
				characters = characters[1:]
			}
		case "R", "N", "-", " ":
			count := 1
			for count < len(characters) && characters[count] == characters[0] {
				count++
			}
			formatted += strings.Repeat(characters[0], count)
			characters = characters[count:]
		default:
			formatted = "YY-RRRRR"
			characters = []string{}
		}
	}
	f.Format = formatted
	uadmin.Save(f)
}

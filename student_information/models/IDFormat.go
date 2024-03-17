package models

import (
	"fmt"
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
		for i := 0; i < len(characters); i++ {
			if characters[i] == "Y" {
				if i+3 < len(characters) && characters[i+1] == "Y" && characters[i+2] == "Y" && characters[i+3] == "Y" {
					formatted += "YYYY"
					characters = append(characters[:i], characters[i+4:]...)
					break
				} else if i+2 < len(characters) && characters[i+1] == "Y" && characters[i+2] == "Y" {
					formatted += "YY"
					characters = append(characters[:i], characters[i+3:]...)
					break
				} else if i+1 < len(characters) && characters[i+1] == "Y" {
					formatted += "YY"
					characters = append(characters[:i], characters[i+2:]...)
					break
				}
				characters = append(characters[:i], characters[i+1:]...)
				break
			} else if characters[i] == "C" {
				if i+1 < len(characters) && characters[i+1] == "C" {
					formatted += "CC"
					characters = append(characters[:i], characters[i+2:]...)
					break
				}
				characters = append(characters[:i], characters[i+1:]...)
				break
			} else if characters[i] == "S" {
				if i+1 < len(characters) && characters[i+1] == "S" {
					formatted += "SS"
					characters = append(characters[:i], characters[i+2:]...)
					break
				}
				characters = append(characters[:i], characters[i+1:]...)
				break
			} else if characters[i] == "R" {
				rCount := 1
				for i+rCount < len(characters) && characters[i+rCount] == "R" {
					rCount++
				}
				formatted += fmt.Sprintf("%s", strings.Repeat("R", rCount))
				characters = append(characters[:i], characters[i+rCount:]...)
				break
			} else if characters[i] == "N" {
				rCount := 1
				for i+rCount < len(characters) && characters[i+rCount] == "N" {
					rCount++
				}
				formatted += fmt.Sprintf("%s", strings.Repeat("N", rCount))
				characters = append(characters[:i], characters[i+rCount:]...)
				break
			} else if characters[i] == "-" {
				rCount := 1
				for i+rCount < len(characters) && characters[i+rCount] == "-" {
					rCount++
				}
				formatted += fmt.Sprintf("%s", strings.Repeat("-", rCount))
				characters = append(characters[:i], characters[i+rCount:]...)
				break
			} else if characters[i] == " " {
				rCount := 1
				for i+rCount < len(characters) && characters[i+rCount] == " " {
					rCount++
				}
				formatted += fmt.Sprintf("%s", strings.Repeat(" ", rCount))
				characters = append(characters[:i], characters[i+rCount:]...)
				break
			} else {
				// If the character is not in the specified formats, add a default format
				formatted = "YY-RRRRR"
				characters = []string{}
				break
			}
		}
	}
	f.Format = formatted
	uadmin.Save(f)
}

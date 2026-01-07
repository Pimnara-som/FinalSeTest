package entity

import (
	govalidator "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string  `valid:required~title must be 3-100,stringlength(3|100)`
	Price float64 `valid:range(50|5000)~Price must be between 50 and 5000`
	Code  string  `valid:code~Code must start with BK 
	followed by 6 digits (0-9),matches(~[BK][0-9]{6}$)`
}

func ValidateBook(p *Book) (bool, error) {
	ok, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return ok, nil
}

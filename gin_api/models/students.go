package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"CPF" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"RG" validate:"len=9, regexp=^[0-9]*$"`
}

var Students []Student

func ValidatesStudentsData(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

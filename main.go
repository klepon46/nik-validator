package main

import (
	"github.com/klepon46/nik-validator/pkg"
)

func main() {

	validator := pkg.NewValidator()
	err := validator.Validate("1234567890123456")
	if err != nil {
		return
	}

}

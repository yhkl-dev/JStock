package validators

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var jstockValidator *validator.Validate

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		jstockValidator = v
	} else {
		log.Fatal("error validator")
	}
}

func registerValidation(tag string, fn validator.Func) {
	err := jstockValidator.RegisterValidation(tag, fn)
	if err != nil {
		log.Fatal(fmt.Sprintf("validator %s error", tag))
	}
}

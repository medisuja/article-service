package helpers

import (
	//"fmt"

	"reflect"
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
)

func ValidationError(s interface{}, err error) map[string]string {
	switch err.(type) {
	case govalidator.Errors:
		// Use reflect to get the raw struct element
		typ := reflect.TypeOf(s).Elem()
		if typ.Kind() != reflect.Struct {
			return nil
		}

		// This is will contain the errors we return back to user
		errs := map[string]string{}
		// Errors found by the validator
		errsByField := govalidator.ErrorsByField(err.(govalidator.Errors))
		// Loop over our struct fields
		for i := 0; i < typ.NumField(); i++ {
			// Get the field
			f := typ.Field(i)
			// Do we have an error for the field
			e, ok := errsByField[f.Name]

			if ok {
				// Try and get the `json` struct tag
				name := strings.Split(f.Tag.Get("json"), ",")[0]
				// If the name is - we should ignore the field
				if name == "-" {
					continue
				}
				// If the name is not blank we add it our error map
				if name != "" {
					errs[name] = e
					continue
				}
				// Finall if all else has failed just add the raw field name to the
				// error map
				errs[CamelCaseToSnakeCase(f.Name)] = e
			}
		}

		// Return the validation error
		return errs
	}

	return nil
}

func CamelCaseToSnakeCase(camelCase string) (inputUnderScoreStr string) {
	//snake_case to camelCase
	for k, v := range camelCase {
		if isUpperCase(string(v)) && k > 0 && unicode.IsLetter(v) {
			inputUnderScoreStr += "_"
		}
		inputUnderScoreStr += strings.ToLower(string(v))
	}
	return inputUnderScoreStr
}

func isUpperCase(str string) bool {
	if str == strings.ToUpper(str) {
		return true
	}
	return false
}

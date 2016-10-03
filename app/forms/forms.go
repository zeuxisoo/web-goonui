package forms

import (
    "fmt"
    "reflect"
    "strings"

    "github.com/go-macaron/binding"
)

type Form interface {
    binding.Validator
}

func getValidatorCondition(field reflect.StructField, prefix string) string {
    for _, rule := range strings.Split(field.Tag.Get("binding"), ";") {
        if strings.HasPrefix(rule, prefix) {
            return rule[len(prefix) : len(rule)-1]
        }
    }

    return ""
}

func getValidatorMaxSize(field reflect.StructField) string {
    return getValidatorCondition(field, "MaxSize(")
}

func validate(errs binding.Errors, data map[string]interface{}, form Form) binding.Errors {
    data["HasError"]     = len(errs) > 0
    data["ErrorMessage"] = ""

    for _, err := range errs {
        fmt.Printf("\n\n%#v\n\n", err)

        fieldName      := err.FieldNames[0]
        classification := err.Classification

        switch classification {
            case binding.ERR_REQUIRED:
                data["ErrorMessage"] = fmt.Sprintf("The %s is required", fieldName)
                break
            case binding.ERR_MAX_SIZE:
                data["ErrorMessage"] = fmt.Sprintf("The %s too long", fieldName)
            default:
                data["ErrorMessage"] = fmt.Sprintf("Unknown error on %s field with error %s", fieldName, classification)
        }

        return errs
    }

    return errs
}

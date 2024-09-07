package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

// Valid returns true if the FieldErrors map doesn't contain entries.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldError adds an error message to the FieldErrors map (as long as the
// given key is free).
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank returns true if a value is not an empty string.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars returns true if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt returns true if the value is in the permittedValues list.
func PermittedInt(value int, permittedValues ...int) bool {
	for _, permittedValue := range permittedValues {
		if value == permittedValue {
			return true
		}
	}
	return false
}

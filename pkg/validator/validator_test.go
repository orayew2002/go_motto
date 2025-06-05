package validator

import (
	"testing"
)

type VTestStruct struct {
	Name  string `required:"true"`
	Email string `required:"true"`
	Age   int
}

func TestValidate(t *testing.T) {
	t.Run("valid struct", func(t *testing.T) {
		v := VTestStruct{Name: "John", Email: "john@example.com"}

		if err := ValidateRequiredFields(v); err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	})

	t.Run("missing name", func(t *testing.T) {
		v := VTestStruct{Name: "", Email: "john@example.com"}

		if err := ValidateRequiredFields(v); err == nil || err.Error() != "field Name is required" {
			t.Errorf("expected error for missing name, got: %v", err)
		}
	})

	t.Run("missing email", func(t *testing.T) {
		v := VTestStruct{Name: "John", Email: ""}

		if err := ValidateRequiredFields(v); err == nil || err.Error() != "field Email is required" {
			t.Errorf("expected error for missing email, got: %v", err)
		}
	})
}

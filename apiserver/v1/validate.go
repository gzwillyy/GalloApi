package v1

import (
	"github.com/gzwillyy/components/pkg/validation"
	"github.com/gzwillyy/components/pkg/validation/field"
)

// Validate validates that a user object is valid.
func (u *Admin) Validate() field.ErrorList {
	val := validation.NewValidator(u)
	allErrs := val.Validate()

	if err := validation.IsValidPassword(u.Password); err != nil {
		allErrs = append(allErrs, field.Invalid(field.NewPath("password"), err.Error(), ""))
	}

	return allErrs
}

// ValidateUpdate validates that a user object is valid when update.
// Like User.Validate but not validate password.
func (u *Admin) ValidateUpdate() field.ErrorList {
	val := validation.NewValidator(u)
	allErrs := val.Validate()

	return allErrs
}

// other validate

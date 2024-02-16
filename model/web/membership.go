package web

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type SignInRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}

type MemberCreateOrUpdateRequest struct {
	Name      string `validate:"required" json:"name"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
	Phone     string `validate:"required,e164" json:"phone"`
	BirthDate string `validate:"required,datetime=2006-01-02" json:"birth_date"`
	Address   string `json:"address"`
}

func containSpecialCharacter(str string, specialCharacter []string) bool {
	for _, sc := range specialCharacter {
		if strings.Contains(str, sc) {
			return true
		}
	}

	return false
}

func PasswordValidator(password string) error {
	validate := validator.New()

	err := validate.Var(password, "containsany=abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		return errors.New("password must contain lowercase")
	}

	err = validate.Var(password, "containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		return errors.New("password must contain uppercase")
	}

	listSpecialCharacter := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "'", "\"", "/", "%", "?", "\n", "<", ">"}
	symbol := containSpecialCharacter(password, listSpecialCharacter)
	if !symbol {
		return errors.New("password must contain special character")
	}

	return nil
}

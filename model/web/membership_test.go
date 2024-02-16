package web

import (
	"fmt"
	"testing"
)

func TestPasswordValidator(t *testing.T) {
	pass := "123aB!@#"
	err := PasswordValidator(pass)
	fmt.Println(err)
}

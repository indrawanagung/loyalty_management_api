package util

import (
	"fmt"
	"testing"
)

func TestGenerateIDTransaction(t *testing.T) {
	transaction := GenerateIDTransaction(1)
	fmt.Println(transaction)
}
func TestGenerateIDGetMember(t *testing.T) {
	transaction := GenerateIDGetMember(1)
	fmt.Println(transaction)
}

func TestGenerateIDActivityMember(t *testing.T) {
	transaction := GenerateIDActivityMember(1)
	fmt.Println(transaction)
}

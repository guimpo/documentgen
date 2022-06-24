package documentgen

import (
	"fmt"
	"regexp"
	"testing"
)

// Must return 11 characters
func TestCPF(t *testing.T) {
	want := regexp.MustCompile(`\d{11}`)
	cpf, err := CPF(1)
	fmt.Println(cpf)

	if !want.MatchString(cpf) || err != nil {
		t.Fatalf(`CPF = %q, %v, want match for %#q, nil`, cpf, err, want)
	}
}

// Must return the same 11 characters
func TestCPFSeed(t *testing.T) {
	cpf1, err1 := CPF(1)
	fmt.Println(cpf1)
	cpf2, err2 := CPF(1)
	fmt.Println(cpf2)

	if cpf1 != cpf2 || err1 != nil || err2 != nil {
		t.Fatalf("CPF is not equal")
	}
}

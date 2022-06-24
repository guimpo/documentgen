package documentgen

import (
	"fmt"
	"regexp"
	"testing"
)

// Must return 14 characters
func TestCNPJ(t *testing.T) {
	want := regexp.MustCompile(`\d{14}`)
	cnpj, err := CNPJ(1)
	fmt.Println(cnpj)

	if !want.MatchString(cnpj) || err != nil {
		t.Fatalf(`CPF = %q, %v, want match for %#q, nil`, cnpj, err, want)
	}
}

// Must return the same 14 characters
func TestCNPJSeed(t *testing.T) {
	cnpj1, err1 := CNPJ(1)
	fmt.Println(cnpj1)
	cnpj2, err2 := CNPJ(1)
	fmt.Println(cnpj2)

	if cnpj1 != cnpj2 || err1 != nil || err2 != nil {
		t.Fatalf("CNPJ is not equal")
	}
}

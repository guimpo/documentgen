package documentgen

import (
	"fmt"
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCPF(t *testing.T) {
	want := regexp.MustCompile(`\d{11}`)
	cpf, err := CPF(1)
	fmt.Println(cpf)
	if !want.MatchString(cpf) || err != nil {
		t.Fatalf(`CPF = %q, %v, want match for %#q, nil`, cpf, err, want)
	}
}

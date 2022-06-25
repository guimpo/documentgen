package documentgen

import (
	"fmt"
	"math/rand"
	"strings"
)

func CNPJ(seed int64) (string, error) {
	rand.Seed(seed)

	const cnpjStDigitIdx = 1
	const cnpjNdDigitIdx = 0
	cnpjMultiplicationFactorI := [14]int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}
	cnpjMultiplicationFactorII := [14]int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6}

	generatedCNPJ := [14]int{-1, -1, 1, 0, 0, 0, -1, -1, -1, -1, -1, -1, -1, -1}
	sum := 0

	for i := 2; i < 14; i++ {
		num := 0
		if i <= 5 {
			num = generatedCNPJ[i]
		} else {
			num = rand.Intn(10)
			generatedCNPJ[i] = num
		}
		sum += num * cnpjMultiplicationFactorI[i-2]
	}

	reminder := sum % 11
	result := 11 - reminder
	digit := 0

	if result < 10 {
		digit = result
	}

	generatedCNPJ[cnpjStDigitIdx] = digit

	sum = 0
	for i := 1; i < 14; i++ {
		sum += generatedCNPJ[i] * cnpjMultiplicationFactorII[i-1]
	}

	reminder = sum % 11
	result = 11 - reminder
	digit = 0
	if result < 10 {
		digit = result
	}

	generatedCNPJ[cnpjNdDigitIdx] = digit

	var cpfString strings.Builder
	for d := len(generatedCNPJ) - 1; d >= 0; d-- {
		digitString := fmt.Sprint(generatedCNPJ[d])
		cpfString.WriteString(digitString)
	}
	return cpfString.String(), nil
}

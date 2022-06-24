package documentgen

import (
	"fmt"
	"math/rand"
	"strings"
)

func CNPJ(seed int64) (string, error) {
	rand.Seed(seed)

	const CNPJ_ST_DIGIT_INDEX = 1
	const CNPJ_ND_DIGIT_INDEX = 0
	CNPJ_MULTIPLICATION_FACTOR_1 := [14]int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}
	CNPJ_MULTIPLICATION_FACTOR_2 := [14]int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6}

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
		sum += num * CNPJ_MULTIPLICATION_FACTOR_1[i-2]
	}

	reminder := sum % 11
	result := 11 - reminder
	digit := 0

	if result < 10 {
		digit = result
	}

	generatedCNPJ[CNPJ_ST_DIGIT_INDEX] = digit

	sum = 0
	for i := 1; i < 14; i++ {
		sum += generatedCNPJ[i] * CNPJ_MULTIPLICATION_FACTOR_2[i-1]
	}

	reminder = sum % 11
	result = 11 - reminder
	digit = 0
	if result < 10 {
		digit = result
	}

	generatedCNPJ[CNPJ_ND_DIGIT_INDEX] = digit

	var cpfString strings.Builder
	for d := len(generatedCNPJ) - 1; d >= 0; d-- {
		digitString := fmt.Sprint(generatedCNPJ[d])
		cpfString.WriteString(digitString)
	}
	return cpfString.String(), nil
}

package documentgen

import (
	"fmt"
	"math/rand"
	"strings"
)

func CPF(seed int64) (string, error) {
	rand.Seed(seed)
	const TOP_NINE = 9

	generatedCPJ := [11]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	const TOP_TEN = 10
	const CPF_ST_DIGIT_INDEX = 9  // 10 - 1
	const CPF_ND_DIGIT_INDEX = 10 //  11 - 1

	CPF_MULTIPLICATION_FACTOR_1 := [11]int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	CPF_MULTIPLICATION_FACTOR_2 := [11]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0

	for i := 0; i < TOP_NINE; i++ {
		num := rand.Intn(10)
		generatedCPJ[i] = num
		sum += num * CPF_MULTIPLICATION_FACTOR_1[i]
	}

	reminder := sum % 11
	result := 11 - reminder
	digit := 0

	if result < 10 {
		digit = result
	}

	generatedCPJ[CPF_ST_DIGIT_INDEX] = digit

	sum = 0
	for i := 0; i < TOP_TEN; i++ {
		sum += generatedCPJ[i] * CPF_MULTIPLICATION_FACTOR_2[i]
	}

	reminder = sum % 11
	result = 11 - reminder
	digit = 0

	if result < 10 {
		digit = result
	}

	generatedCPJ[CPF_ND_DIGIT_INDEX] = digit

	var cpfString strings.Builder

	for d := 0; d < len(generatedCPJ); d++ {
		digitString := fmt.Sprint(generatedCPJ[d])
		cpfString.WriteString(digitString)
	}

	return cpfString.String(), nil
}

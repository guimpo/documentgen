package documentgen

import (
	"fmt"
	"math/rand"
	"strings"
)

func CPF(seed int64) (string, error) {
	rand.Seed(seed)
	const topNine = 9

	generatedCPJ := [11]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	const topTen = 10
	const cpfStDigitIdx = 9  // 10 - 1
	const cpfNdDigitIdx = 10 //  11 - 1

	cpfMultiplicationFactorI := [11]int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfMultiplicationFactorII := [11]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0

	for i := 0; i < topNine; i++ {
		num := rand.Intn(10)
		generatedCPJ[i] = num
		sum += num * cpfMultiplicationFactorI[i]
	}

	reminder := sum % 11
	result := 11 - reminder
	digit := 0

	if result < 10 {
		digit = result
	}

	generatedCPJ[cpfStDigitIdx] = digit

	sum = 0
	for i := 0; i < topTen; i++ {
		sum += generatedCPJ[i] * cpfMultiplicationFactorII[i]
	}

	reminder = sum % 11
	result = 11 - reminder
	digit = 0

	if result < 10 {
		digit = result
	}

	generatedCPJ[cpfNdDigitIdx] = digit

	var cpfString strings.Builder

	for d := 0; d < len(generatedCPJ); d++ {
		digitString := fmt.Sprint(generatedCPJ[d])
		cpfString.WriteString(digitString)
	}

	return cpfString.String(), nil
}

package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringToFloat(numString string) (float64, error) {
	num, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return 0, err
	}

	formattedFloat := fmt.Sprintf("%.2f", num)

	result, err := strconv.ParseFloat(formattedFloat, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GeneratePastYears(amount int) []int {
	currentYear := time.Now().Year()
	years := make([]int, 0)

	for i := 0; i < amount; i++ {
		years = append(years, currentYear-i)
	}

	return years
}

func GenerateYearsUntil(startYear int) []string {
	currentYear := time.Now().Year()

	years := make([]string, 0)

	for year := currentYear; year >= startYear; year-- {
		years = append(years, fmt.Sprintf("%d", year))
	}

	return years
}

func GetLastSentence(sentance string) string {
	words := strings.Split(sentance, " ")
	lastWord := words[len(words)-1]
	return lastWord
}

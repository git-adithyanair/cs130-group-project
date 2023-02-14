package util

import (
	"math/rand"
	"strings"
)

const digits = "0123456789"
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int, data string) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		randInt := rand.Intn(len(data))
		sb.WriteByte(data[randInt])
	}

	return sb.String()
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomNumericString(length int) string {
	return randomString(length, digits)
}

func RandomString(length int) string {
	return randomString(length, alphabet)
}

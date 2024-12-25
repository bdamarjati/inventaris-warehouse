package util

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

var categories = [...]string{"Stationery", "Tools", "Packaging", "Safety"}
var statuses = [...]string{"Ready", "Procured", "Arrived"}

func randomString(n int) string {
	text := make([]byte, n)
	for i := range text {
		text[i] = letters[rand.Intn(len(letters))]
	}
	return string(text)
}

func RandomName() string {
	return randomString(6)
}

func RandomText(n int) string {
	return randomString(n)
}

func RandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

func RandomCategory() string {
	return categories[rand.Intn(len(categories))]
}

func RandomStatus() string {
	return statuses[rand.Intn(len(statuses))]
}

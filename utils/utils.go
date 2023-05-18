package utils

import "math/rand"

const ChararcterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ShortPrefix = "http://sho.rt/"
const DbKey = "DB"

var Rand rand.Source

func GenerateRandomCharacters(chararcterSet string, numberOfDigits int) string {
	bytes := make([]byte, numberOfDigits)
	for k := range bytes {
		bytes[k] = chararcterSet[rand.Intn(len(chararcterSet))]
	}
	return string(bytes)
}

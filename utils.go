package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func calculateHash(input string) string {
	data := []byte(input)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}

func cleanString(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}

func nameSlug(firstName string, lastName string) string {
	slug := firstName + "-" + lastName
	slug = cleanString(slug)

	return strings.ToLower(slug)
}

func GenerateGameId(year, gameNumber int) string {
	// yeargameNumber(00-padded)
	return fmt.Sprintf("%d%03d", year, gameNumber)
}

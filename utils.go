package domain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"110yards.ca/libs/go/core/logger"
)

func CalculateHash(o interface{}) (string, error) {
	// marshal p to json
	j, err := json.Marshal(o)

	if err != nil {
		logger.Errorf("Error while marshalling player: %s", err.Error())
		return "", err
	}

	hash := HashString(string(j))

	return hash, nil
}

func HashString(s string) string {
	data := []byte(s)
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

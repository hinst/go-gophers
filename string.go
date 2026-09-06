package gophers

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/text/unicode/norm"
)

func GetIntFromString(text string) int {
	return AssertResultError(strconv.Atoi(text))
}

func GetInt32FromString(text string) int32 {
	return int32(AssertResultError(strconv.ParseInt(text, 10, 32)))
}

func GetInt64FromString(text string) int64 {
	return AssertResultError(strconv.ParseInt(text, 10, 64))
}

func GetInt64FromStringOptional(text string) *int64 {
	if text == "" {
		return nil
	}
	var value = GetInt64FromString(text)
	return &value
}

func GetStringFromInt64(number int64) string {
	return strconv.FormatInt(number, 10)
}

func GetStringFromInt(number int) string {
	return strconv.Itoa(number)
}

func GetStringFromBool(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func GetQuotedString(text string) string {
	return "\"" + text + "\""
}

func RequireEnvVar(name string) string {
	var value = os.Getenv(name)
	if value == "" {
		log.Fatalln("Environment variable is required:", name)
	}
	return value
}

func ReadEnvVar(name string, defaultValue string) string {
	var value = os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func NormalizeString(text string) string {
	return norm.NFC.String(text)
}

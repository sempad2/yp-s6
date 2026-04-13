package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Converter(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("empty string")
	}

	// допустимые символы для Морзе
	allowed := ".- \n\t/"

	for _, char := range s {
		if !strings.ContainsRune(allowed, char) {
			return morse.ToMorse(s), nil
		}
	}

	return morse.ToText(s), nil
}

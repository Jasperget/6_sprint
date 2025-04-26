package service

import (
	"6_sprint/pkg/morse"
	"errors"
	"strings"
)

// ConvertTextToMorseOrViceVersa определяет тип входной строки (код Морзе или текст)
// и выполняет соответствующее преобразование.
func ConvertTextToMorseOrViceVersa(input string) (string, error) {
	// Проверяем, является ли строка кодом Морзе
	if isMorseCode(input) {
		// Конвертируем код Морзе в текст
		return morse.ToText(input), nil
	} else if isPlainText(input) {
		// Конвертируем текст в код Морзе
		return morse.ToMorse(input), nil
	}
	// Если строка не соответствует ни одному формату, возвращаем ошибку
	return "", errors.New("invalid input format: neither Morse code nor plain text")
}

// isMorseCode проверяет, является ли строка кодом Морзе.
func isMorseCode(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

// isPlainText проверяет, является ли строка обычным текстом.
func isPlainText(input string) bool {
	return strings.TrimSpace(input) != ""
}

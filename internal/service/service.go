package service

import (
	"strings"

	"6_sprint/pkg/morse"
)

// ConvertTextToMorseOrViceVersa определяет тип входной строки (код Морзе или текст)
// и выполняет соответствующее преобразование.
func ConvertTextToMorseOrViceVersa(input string) (string, error) {
	// Проверяем, является ли строка кодом Морзе
	if isMorseCode(input) {
		// Конвертируем код Морзе в текст
		return morse.ToText(input), nil
	}

	// Проверяем, является ли строка обычным текстом
	if isPlainText(input) {
		// Конвертируем текст в код Морзе
		return morse.ToMorse(input), nil
	}

	// Если строка не соответствует ни одному формату, возвращаем ошибку
	return "", morse.ErrNoEncoding{Text: input}
}

// isMorseCode проверяет, является ли строка кодом Морзе.
func isMorseCode(input string) bool {
	return !strings.ContainsFunc(input, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
}

// isPlainText проверяет, является ли строка обычным текстом.
func isPlainText(input string) bool {
	return strings.TrimSpace(input) != ""
}

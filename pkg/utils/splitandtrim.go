package utils

import "strings"

func SplitAndTrimSpaces(str string) []string {
	// Dividir a string em um slice de strings usando a função Split
	arrayDeStrings := strings.Split(str, ",")

	// Remover espaços no início e no final de cada string no array
	for i := range arrayDeStrings {
		arrayDeStrings[i] = strings.TrimSpace(arrayDeStrings[i])
	}

	return arrayDeStrings
}

package util

import (
	"encoding/json"
	"fmt"
)

func PrintStructColor(title string, v any) {
	fmt.Printf("\033[34m%s:\033[0m\n", title) // título en azul

	if v == nil {
		return
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("\033[31mError al imprimir %s: %v\033[0m\n", title, err)
		return
	}

	fmt.Printf("\033[32m%s\033[0m\n", string(data)) // JSON en verde
}

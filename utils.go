package main

import (
	"bufio"
	"fmt"
	"os"
)

func getUserInput(requiredParams []string) map[string]string {
	inputs := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for _, param := range requiredParams {
		fmt.Printf("Enter %s: ", param)
		scanner.Scan()
		inputs[param] = scanner.Text()
	}

	return inputs
}

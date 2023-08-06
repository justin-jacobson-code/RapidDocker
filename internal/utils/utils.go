package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/justin-jacobson-code/RapidDocker/internal/dbconfig"
)

func GetUserInput(config *dbconfig.Configuration) {
	scanner := bufio.NewScanner(os.Stdin)

	params := config.Params

	fmt.Printf("Would you like to setup advanced options for %s? (y/n): ", config.ImageName)
	scanner.Scan()

	if strings.ToLower(scanner.Text()) == "y" {
		config.UseAdvancedOptions = true
	} else {
		config.UseAdvancedOptions = false
	}

	if config.UseAdvancedOptions {
		for key, value := range params {
			fmt.Printf("Enter %s (%s): ", key, value)
			scanner.Scan()

			input := scanner.Text()
			// check if input provided for username or password
			// if no input provided, ask again
			for (key == "username" || key == "password") && len(input) == 0 {
				fmt.Printf("%s is required, please enter: ", key)
				scanner.Scan()
				input = scanner.Text()
			}

			if len(input) > 0 {
				config.Params[key] = input
			}
		}
	}
}

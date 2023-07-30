package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/justinj96/RapidDocker/internal/dbconfig"
)

func GetUserInput(config dbconfig.Configuration) {
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
			if len(scanner.Text()) > 0 {
				config.Params[key] = scanner.Text()
			}
		}
	}
}

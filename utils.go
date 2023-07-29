package main

import (
	"bufio"
	"fmt"
	"os"
)

func getUserInput(config Configuration) {
	scanner := bufio.NewScanner(os.Stdin)

	params := config.Params

fmt.Printf("Would you like to setup the advanced options for %s? (y/n): ", config.ImageName)
	scanner.Scan()

	if scanner.Text() == "y" {
		for key, value := range params {
			fmt.Printf("Enter %s (%s): ", key, value)
			scanner.Scan()
			if len(scanner.Text()) > 0 {
				config.Params[key] = scanner.Text()
			}
		}
	}

	// for _, param := range optionalParams {
	// 	fmt.Printf("Enter %s or leave blank to default: ", param)
	// 	scanner.Scan()
	// 	inputs[param] = scanner.Text()
	// }

	// for _, param := range advancedParams {
	// 	fmt.Printf("Enter %s or leave blank to default: ", param)
	// 	scanner.Scan()
	// 	inputs[param] = scanner.Text()
	// }
}

func MongoConnectionString(config Configuration) string {
	connectionString := config.ConnectionString

	// if user and password are provided, update connection string
	if len(config.Params["user"]) > 0 && len(config.Params["password"]) > 0 {
		connectionString = fmt.Sprintf("mongodb://%s:%s@localhost:%s",
			config.Params["user"],
			config.Params["password"],
			config.Params["port"])
	}
	return connectionString
}

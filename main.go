package main

import (
	"fmt"
	"strings"

	"github.com/justinj96/RapidDocker/internal/dbconfig"
	"github.com/justinj96/RapidDocker/internal/utils"
)

func main() {
	// Prompt the user for the desired container type
	fmt.Println("Available container types:")
	for containerType := range dbconfig.Configurations {
		fmt.Println(containerType)
	}

	fmt.Print("Enter the container type: ")
	var containerType string
	fmt.Scanln(&containerType)

	// Check if the selected container type exists in the configurations map
	config, exists := dbconfig.Configurations[strings.ToLower(containerType)]
	if !exists {
		fmt.Println("Invalid container type.")
		return
	}

	utils.GetUserInput(config)

	// Use the configuration values to create the container
	fmt.Printf("Creating container with image '%s'...\n", config.ImageName)

	cli := utils.CreateDockerClient()

	fmt.Println("Pulling Docker image...")

	utils.PullDockerImage(cli, config.ImageName, config.Params["tag"])

	fmt.Println("Running Docker container...")
	utils.RunDockerContainer(cli, config)

	config.GenerateConnectionString(&config)
	fmt.Println("Connection string:", config.ConnectionString)
	// if len(config.ConnectionString) > 0 {
	// 	fmt.Println("Connection string:", config.ConnectionString)
	// }

	fmt.Printf("Use \"docker rm %s\" to remove container\n", config.Params["containerName"])
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Prompt the user for the desired container type
	fmt.Println("Available container types:")
	for containerType := range Configurations {
		fmt.Println(containerType)
	}

	fmt.Print("Enter the container type: ")
	var containerType string
	fmt.Scanln(&containerType)

	// Check if the selected container type exists in the configurations map
	config, exists := Configurations[strings.ToLower(containerType)]
	if !exists {
		fmt.Println("Invalid container type.")
		return
	}

	// get config values from user
	_ = getUserInput(config.RequiredParams)

	// Use the configuration values to create the container
	fmt.Printf("Creating container with image '%s'...\n", config.ImageName)

	// cli := createDockerClient()

	// fmt.Println("Pulling Docker image...")

	// pullDockerImage(cli, configValues.ImageName)

	// fmt.Println("Running Docker container...")
	// runDockerContainer(cli, imageName)

	// Perform tasks with the running container here...

	// fmt.Println("Stopping and removing Docker container...")
	// err = stopAndRemoveContainer(cli, containerID)
	// if err != nil {
	// 	fmt.Println("Error stopping and removing Docker container:", err)
	// 	return
	// }
}

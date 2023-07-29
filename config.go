package main

// Configuration represents the settings required for a specific container type.
type Configuration struct {
	ImageName      string
	RequiredParams []string
}

// Configurations is a map that stores the configurations for different container types.
var Configurations = map[string]Configuration{
	"mongodb": {
		ImageName: "mongo:latest",
		RequiredParams: []string{
			"user",
			"password",
		},
	},
	"postgresql": {
		ImageName: "postgresql",
		RequiredParams: []string{
			"user",
			"password",
		},
	},
	// Add other container configurations here...
}

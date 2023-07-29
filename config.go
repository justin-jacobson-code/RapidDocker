package main

// Configuration represents the settings required for a specific container type.
type Configuration struct {
	ImageName                string
	HostIp                   string
	Params                   map[string]string
	ConnectionString         string                            // optional
	GenerateConnectionString func(config Configuration) string // optional
}

// Configurations is a map that stores the configurations for different container types.
var Configurations = map[string]Configuration{
	"mongodb": {
		ImageName: "mongo",
		HostIp:    "0.0.0.0",
		Params: map[string]string{
			"port":          "27017",
			"user":          "",
			"password":      "",
			"database":      "go",
			"containerName": "mongo-go",
			"tag":           "latest",
		},
		ConnectionString:         "mongodb://localhost:27017",
		GenerateConnectionString: MongoConnectionString,
	},
}

package dbconfig

// Configuration represents the settings required for a specific container type.
type Configuration struct {
	ImageName            string
	HostIp               string
	UseAdvancedOptions   bool
	Params               map[string]string
	EnvironmentVariables []string // optional

	// connection strings are only printed to the console for convienience
	// the program itself does not use them, feel free to leave them out
	ConnectionString         string                      // optional
	GenerateConnectionString func(config *Configuration) // optional
	EnvVarSetup              func(config *Configuration) // optional
}

// Configurations is a map that stores the configurations for different container types.
var Configurations = map[string]Configuration{
	"mongodb": {
		ImageName:          "mongo",
		HostIp:             "0.0.0.0",
		UseAdvancedOptions: false,

		Params: map[string]string{
			"port":          "27017",
			"username":      "",
			"password":      "",
			"containerName": "mongo-rapid-docker",
			"tag":           "latest",
		},

		EnvironmentVariables:     []string{},
		ConnectionString:         "mongodb://localhost:27017",
		GenerateConnectionString: mongoConnectionString,
		EnvVarSetup:              mongoEnvVarSetup,
	},

	"postgres": {
		ImageName:          "postgres",
		HostIp:             "0.0.0.0",
		UseAdvancedOptions: false,

		Params: map[string]string{
			"port":          "5432",
			"username":      "postgres",
			"password":      "",
			"containerName": "postgres-rapid-docker",
			"tag":           "latest",
		},

		EnvironmentVariables:     []string{},
		ConnectionString:         "postgres://postgres@localhost:5432",
		GenerateConnectionString: postgresConnectionString,
		EnvVarSetup:              postgresEnvVarSetup,
	},
}

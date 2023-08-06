package dbconfig

import (
	"fmt"
)

func mongoConnectionString(config *Configuration) {
	connectionString := config.ConnectionString

	// if user and password are provided, update connection string
	if len(config.Params["username"]) > 0 && len(config.Params["password"]) > 0 {
		connectionString = fmt.Sprintf("mongodb://%s:%s@localhost:%s",
			config.Params["username"],
			"******",
			config.Params["port"])
	}
	config.ConnectionString = connectionString
}

func mongoEnvVarSetup(config *Configuration) {
	if config.UseAdvancedOptions {
		// if user and password are provided, update environment variables
		config.EnvironmentVariables = []string{
			"MONGO_INITDB_ROOT_USERNAME=" + config.Params["username"],
			"MONGO_INITDB_ROOT_PASSWORD=" + config.Params["password"],
		}
	}
}

func postgresConnectionString(config *Configuration) {
	connectionString := config.ConnectionString

	// if user and password are provided, update connection string
	if len(config.Params["username"]) > 0 && len(config.Params["password"]) > 0 {
		connectionString = fmt.Sprintf("postgres://%s:%s@localhost:%s",
			config.Params["username"],
			"******",
			config.Params["port"])
	}
	config.ConnectionString = connectionString
}

func postgresEnvVarSetup(config *Configuration) {

	if config.UseAdvancedOptions {
		// if user and password are provided, update environment variables
		config.EnvironmentVariables = []string{
			"POSTGRES_USER=" + config.Params["username"],
			"POSTGRES_PASSWORD=" + config.Params["password"],
		}
	} else {
		// set postgres host auth method to trust, no password needed
		config.EnvironmentVariables = []string{
			"POSTGRES_HOST_AUTH_METHOD=trust",
		}
	}

}

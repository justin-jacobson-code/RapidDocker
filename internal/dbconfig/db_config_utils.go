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
			"**********",
			config.Params["port"])
	}
	config.ConnectionString = connectionString
}

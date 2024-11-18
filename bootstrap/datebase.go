package bootstrap

import database "h5/pkg/database"

func SetupDatabase() {
	database.Initialize()
}

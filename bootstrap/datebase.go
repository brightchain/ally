package bootstrap

import database "ally/pkg/database"

func SetupDatabase() {
	database.Initialize()
}

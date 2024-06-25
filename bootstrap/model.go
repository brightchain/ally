package bootstrap

import "ally/pkg/model"

func SetupModel() {
	model.InitDb()
}

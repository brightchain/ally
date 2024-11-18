package bootstrap

import "h5/pkg/model"

func SetupModel() {
	model.InitDb()
}

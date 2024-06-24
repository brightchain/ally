package bootstrap

import "ally/pkg/viperConf"

func SetupViper() {
	viperConf.Initialize()
	go viperConf.DynamicReloadConfig()
}

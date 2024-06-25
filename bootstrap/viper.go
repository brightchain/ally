package bootstrap

import "ally/pkg/config"

func SetupViper() {
	config.Initialize()
	//go viperConf.DynamicReloadConfig()
}

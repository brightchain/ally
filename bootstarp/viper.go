package bootstarp

import "ally/pkg/viperConf"

func SetupViper() {
	viperConf.Initialize()
	go viperConf.DynamicReloadConfig()
}

package viperConf

import (
	"os"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

var Data *viper.Viper

func Initialize() *viper.Viper {
	env := os.Getenv("GIN_MODE")
	name := "config"
	if env != "" {
		name += "_" + env
	}
	Data = viper.New()
	Data.SetConfigName(name)
	Data.SetConfigType("yaml")
	Data.AddConfigPath(".")

	err := Data.ReadInConfig()
	if err != nil {
		slog.Error("配置文件读取失败！", err)
	}

	return Data
}

func DynamicReloadConfig() {
	Data.WatchConfig()
}

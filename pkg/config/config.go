package config

type AppConfig struct {
	App    App    `json:"app"`
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
	Logger Logger `json:"logger"`
}

type App struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	DbName   string `json:"dbName"`
}

type Redis struct {
	Host     string `json:"host"`
	Db       int    `json:"db"`
	Port     int    `json:"port"`
	Password string `json:"Password"`
}

type Logger struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
	MaxSize    int    `json:"maxSize"`
}

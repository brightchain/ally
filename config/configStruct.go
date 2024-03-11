package config

type AppConfig struct {
	App   App   `json:"app"`
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type App struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Db       string `json:"db"`
}

type Redis struct {
	Host string `json:"host"`
	Db   string `json:"db"`
	Port int    `json:"port"`
}

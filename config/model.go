package config

type IConfig interface {
	Get() *Config
}

type IConfigModel struct {
	model *Config
}

type Config struct {
	Server   Server   `json:"server"`
	Postgres Postgres `json:"postgres"`
}

type Server struct {
	Http Http `json:"http"`
}

type Http struct {
	Address string `json:"address"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

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
	Mongodb  Mongodb  `json:"mongodb"`
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

type Mongodb struct {
	Hosts       []string `json:"hosts"`
	RetryWrites bool     `json:"retryWrites"`
	ReplicaSet  string   `json:"replicaSet"`
	AppName     string   `json:"appName"`
	Uri         string   `json:"uri"`
	Db          string   `json:"db"`
	Collection  string   `json:"collection"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	AuthSource  string   `json:"authSource"`
	MinPoolSize uint64   `json:"minPoolSize"`
	MaxPoolSize uint64   `json:"maxPoolSize"`
}

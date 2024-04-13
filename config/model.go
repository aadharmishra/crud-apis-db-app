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
	Redis    Redis    `json:"redis"`
	OAuth    OAuth    `json:"oAuth"`
	Google   Google   `json:"google"`
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

type Redis struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type OAuth struct {
	Web Web `json:"web"`
}

type Web struct {
	Client_id                   string   `json:"client_id"`
	Project_id                  string   `json:"project_id"`
	Auth_uri                    string   `json:"auth_uri"`
	Token_uri                   string   `json:"token_uri"`
	Auth_provider_x509_cert_url string   `json:"auth_provider_x509_cert_url"`
	Client_secret               string   `json:"client_secret"`
	Redirect_uris               []string `json:"redirect_uris"`
	Javascript_origins          []string `json:"javascript_origins"`
	Scopes                      []string `json:"scopes"`
}

type Google struct {
	UserInfoUrl          string `json:"userInfoUrl"`
	YoutubeSearchUrl     string `json:"youtubeSearchUrl"`
	GoogleDriveUploadUrl string `json:"googleDriveUploadUrl"`
}

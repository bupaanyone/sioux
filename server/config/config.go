package config

type Config struct {
	Service Service `json:"service"`
	Handler Handler `json:"handler"`
}

type Service struct {
	DbName string `json:"db_name"`

	PasswordSalt      string `json:"password_salt"`
	PasswordIteration int    `json:"password_iteration"`

	RootUsername string `json:"root_username"`
	RootPassword string `json:"root_password"`
	RootInitFile string `json:"root_init_file"`
}

type Handler struct {
	UrlPrefix string `json:"url_prefix"`
	Host      string `json:"host"`
}

var C Config

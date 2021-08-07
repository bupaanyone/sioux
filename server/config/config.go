package config

type Config struct {
	Handler Handler `json:"handler"`
}

type Handler struct {
	UrlPrefix string `json:"url_prefix"`
	Host      string `json:"host"`
}

var C Config

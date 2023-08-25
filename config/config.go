package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
	Redis    Redis
	Ouath    Oauth
	JWT      JWT
	Firebase Firebase
	Mail     Mail
}

type Server struct {
	Port int
	Env  string
}

type Postgres struct {
	DSN      string
	MaxIdle  int
	MaxConns int
	IdleTime time.Duration
	LifeTime time.Duration
}

type Redis struct {
	DSN  string
	Port int
}

type Oauth struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
	RedirectURL  string
	State        string
}

type JWT struct {
	Secret           string
	ExpiresIn        time.Duration
	RefreshSecret    string
	RefreshExpiresIn time.Duration
}

type Firebase struct {
	Bucket       string
	StorageToken string
	ConfigPath   string
}

type Mail struct {
	Sender string
	Pass   string
}

func LoadConfig(filename string) (config *Config, err error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}

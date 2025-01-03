package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal      = "local"
	EnvProduction = "production"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	TokenKey string         `yaml:"tokenKey" env-default:"secret_key"`
	Port     int32          `yaml:"port" env-default:"8080"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

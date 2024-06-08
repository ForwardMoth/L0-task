package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env              string `yaml:"env" env-default:"local"`
	HttpServerConfig `yaml:"http_server"`
	DBConfig         `yaml:"db"`
}

type HttpServerConfig struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dialect  string `yaml:"dialect"`
}

func MustLoad() *Config {
	configPath := os.Args[1]

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist, %v", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Can't read config file: %v", configPath)
	}

	return &config
}

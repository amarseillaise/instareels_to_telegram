package services

import (
	"time"

	gored "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr        string        `yaml:"addr"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

cfg := Config{
	Addr: "127.0.0.1:6379"
	Password:
}
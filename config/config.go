package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

const path = "./config/config.json"

type Config struct {
	MaxGenerationsPerMonth int64
	Postgres               Postgres
	AI                     AI
	TGBot                  TGBot
	ServerCloseTimeoutMS   time.Duration
}

type Postgres struct {
	Host               string
	Port               string
	DBName             string
	User               string
	Password           string
	MaxOpenConns       int
	MaxIdleConns       int
	MaxConnLifetimeMin time.Duration
}

type AI struct {
	APIKey string
	Host   string
}

type TGBot struct {
	Token     string
	Debug     bool
	TimeoutMS time.Duration
}

var (
	config Config
	once   sync.Once
)

func MustGet() Config {
	once.Do(func() {
		fileCfg, err := os.Open(path)
		if err != nil {
			logrus.Panicf("error loading config file: %v", err)
		}

		defer func() {
			_ = fileCfg.Close()
		}()

		decoder := json.NewDecoder(fileCfg)

		err = decoder.Decode(&config)
		if err != nil {
			logrus.Panicf("error decoding config file: %v", err)
		}
	})

	return config
}

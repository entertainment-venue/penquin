package penquin

import (
	"time"
)

type PQConfig struct {
	StorageType    string
	IdGenerator    string
	MsgExpireHours int

	Server struct {
		Port         string
		DialTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Log struct {
		Level     string
		ExpireDay int
	}

	RedisConfig struct {
		Addr         string
		Password     string
		DialTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		MaxRetries   int
	}
}

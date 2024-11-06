package config

import "time"

type Telegram struct {
	PollTimeout time.Duration `yaml:"poll_timeout"`
}

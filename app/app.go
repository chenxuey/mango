package app

import (
	"mango/config"
)

type AppBase struct {
	Err     error
	ErrCode int
	*config.Logger
}

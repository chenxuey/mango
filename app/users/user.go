package users

import (
	"mango/app"
	"mango/config"
	"fmt"
)

type UserHttp struct {
	app.AppBase
}

func NewUserHttp(logger *config.Logger) *UserHttp {
	appHttp := UserHttp{}
	appHttp.Logger = logger
	return &appHttp
}

func (userHttp *UserHttp) Get() {
	fmt.Println("你请求了 user GET 方法", userHttp.Info)
	userHttp.Println("你请求了 user GET 方法")
}

package config

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func New() *iris.Application {
	app := iris.New()
	return app
}

func Party(app *iris.Application, dir string) router.Party {
	return app.Party(dir)
}

func Listen(app *iris.Application) {
	app.Listen(":8080")
}

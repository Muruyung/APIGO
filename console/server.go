package console

import (
	"APIGO/config"
	"APIGO/controller/admin"
	"APIGO/controller/user"
)

func Run() {
	app := config.New()
	db := config.DBase()
	user.Init(app, db)
	admin.Init(app, db)
	config.Listen(app)
}

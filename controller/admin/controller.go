package admin

import (
	"APIGO/config"
	"APIGO/model"
	"APIGO/model/admin"
	"fmt"

	"github.com/kataras/iris/v12"

	"xorm.io/xorm"
)

var DB *xorm.Engine

func Init(app *iris.Application, db *xorm.Engine) {
	DB = db
	err := db.Sync2(new(model.Tb_admin))

	if err != nil {
		fmt.Println(err)
		return
	}

	adminAPI := config.Party(app, "/admin")
	{
		adminAPI.Use(iris.Compression)
		adminAPI.Get("/", GetAll)
	}
}

func GetAll(ctx iris.Context) {
	ctx.JSON(admin.GetAll(DB))
}

package admin

import (
	"APIGO/config"
	"APIGO/model"
	"APIGO/model/admin"
	"fmt"
	"strconv"

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
		adminAPI.Post("/", Create)
		adminAPI.Get("/", ReadAll)
		adminAPI.Put("/", Update)
		adminAPI.Delete("/{id}", Delete)
	}
}

func Create(ctx iris.Context) {
	var tb_admin model.Tb_admin
	err := ctx.ReadJSON(&tb_admin)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.Post(DB, tb_admin)
}

func ReadAll(ctx iris.Context) {
	ctx.JSON(admin.GetAll(DB))
}

func Update(ctx iris.Context) {
	var tb_admin model.Tb_admin
	err := ctx.ReadJSON(&tb_admin)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.Put(DB, tb_admin)
}

func Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	admin.Delete(DB, id)
}

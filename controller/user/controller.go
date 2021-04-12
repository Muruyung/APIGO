package user

import (
	"APIGO/config"
	"APIGO/model"
	"APIGO/model/user"
	"fmt"
	"strconv"

	"github.com/kataras/iris/v12"

	"xorm.io/xorm"
)

var DB *xorm.Engine

func Init(app *iris.Application, db *xorm.Engine) {
	DB = db
	err := db.Sync2(new(model.Tb_user))

	if err != nil {
		fmt.Println(err)
		return
	}

	userAPI := config.Party(app, "/user")
	{
		userAPI.Use(iris.Compression)
		userAPI.Post("/", Create)
		userAPI.Get("/", ReadAll)
		userAPI.Put("/", Update)
		userAPI.Delete("/{id}", Delete)
	}
}

func Create(ctx iris.Context) {
	var tb_user model.Tb_user
	err := ctx.ReadJSON(&tb_user)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.Post(DB, tb_user)
}

func ReadAll(ctx iris.Context) {
	ctx.JSON(user.GetAll(DB))
}

func Update(ctx iris.Context) {
	var tb_user model.Tb_user
	err := ctx.ReadJSON(&tb_user)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.Put(DB, tb_user)
}

func Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.Delete(DB, id)
}

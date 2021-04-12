package user

import (
	"APIGO/config"
	"APIGO/model"
	"APIGO/model/user"
	response "APIGO/view"
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

	err = user.Post(DB, tb_user)
	response.Show(ctx, err)
}

func ReadAll(ctx iris.Context) {
	res, err := user.GetAll(DB)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(res)
}

func Update(ctx iris.Context) {
	var tb_user model.Tb_user
	err := ctx.ReadJSON(&tb_user)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = user.Put(DB, tb_user)
	response.Show(ctx, err)
}

func Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = user.Delete(DB, id)
	response.Show(ctx, err)
}

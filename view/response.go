package response

import (
	"github.com/kataras/iris/v12"
)

func Show(ctx iris.Context, err error) {
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	} else {
		message := ctx.PostValue("message")
		ctx.JSON(iris.Map{
			"Status":  "Success",
			"Message": message,
		})
		return
	}
}

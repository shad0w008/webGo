package control

import (
	"fmt"
	"github.com/kataras/iris"
	"webGo/utils"
)

func TestLogin(ctx iris.Context) {
	if auth,_:=sess.Start(ctx).GetBoolean("authenticated"); !auth{
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	fmt.Println("cookie",ctx.Request().Cookies())
	ctx.JSON(utils.ResultUtil.Success([2]int{1,2}))
}


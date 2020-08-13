package control

import (
	"github.com/kataras/iris"
	"webGo/utils"
)

func GetAa(ctx iris.Context) {
	//name:=ctx.URLParam("name")
	//fmt.Println()
	//fmt.Println(name)
	//fmt.Println(ctx.Params().Get("Host"))
	ctx.JSON(utils.ResultUtil.Success([2]int{1,2}))
}

func Once(ctx iris.Context)  {
	ctx.JSON(utils.ResultUtil.Success([3]string{"q","b","d"}))

}

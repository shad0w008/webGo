package control

import (
	"fmt"
	"github.com/kataras/iris"
	"webGo/database"
	"webGo/utils"
)

func GetAa(ctx iris.Context) {
	if auth,_:=sess.Start(ctx).GetBoolean("authenticated"); !auth{
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	a,err:=database.Engine.Table("vulzc").Count()
	if err!=nil{
		fmt.Print(a)
	}
	fmt.Println(a)
	ctx.JSON(utils.ResultUtil.Success([2]int{1,2}))
}

func Once(ctx iris.Context)  {
	ctx.JSON(utils.ResultUtil.Success([3]string{"q","b","d"}))

}

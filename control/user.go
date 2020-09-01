package control

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"webGo/model"
	"webGo/service"
	"webGo/utils"
)
var (
	cookieNameForSessionID = "sessionId"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})

)
func Login(ctx iris.Context)  {
	session := sess.Start(ctx)
	session.Set("authenticated",true)
	loginInfo1 := model.UserLogin{}
	err:=ctx.ReadJSON(&loginInfo1)
	if err!=nil{
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	}else {
		aa,_:=service.UserLogin(loginInfo1)
		ctx.JSON(utils.ResultUtil.Success(aa))
	}
}

func LogOut(ctx iris.Context)  {
	session :=sess.Start(ctx)
	session.Set("authenticated",false)
	session.Delete("authenticated")
	session.Destroy()
	ctx.HTML("Logout")

}

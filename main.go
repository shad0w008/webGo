 package main

 import (
	 "fmt"
	 "github.com/kataras/iris"
	 "github.com/kataras/iris/middleware/logger"
	 "github.com/kataras/iris/middleware/recover"
	 "webGo/config"
	 "webGo/control"
 )

 func router(app  *iris.Application)  {

      //app.Get("/", func(ctx iris.Context) {
      //    ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
      //})
	 app.Get("/", func(ctx iris.Context) {
	 	ctx.HTML("<h1>welcole</h1>")

	 })
      apis:=app.Party("/apis")
      {
      	apis.Get("/aa", control.GetAa)
      	apis.Get("/once",control.Once)
	 }

 }
    func main() {
        fmt.Println("aa",config.Conf.Get("app.host").(string))
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        router(app)
        app.Run(iris.Addr(":8086"), iris.WithoutInterruptHandler)
        fmt.Println("aa",config.Conf.Get("app.host").(string))


    }
package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	//1. create iris instance
	app := iris.New()
	//2.setup template
	app.HandleDir("/public", iris.Dir("./frontend/web/public"))
	//3.get aceess to static files
	app.HandleDir("/html", iris.Dir("./frontend/web/htmlProductShow"))
	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

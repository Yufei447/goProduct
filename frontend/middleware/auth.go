package middleware

import "github.com/kataras/iris/v12"

func AuthConProduct(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("Must log in first!")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("Logged in user")
	ctx.Next()
}

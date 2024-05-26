package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/opentracing/opentracing-go/log"
	"go-product/backend/web/controllers"
	"go-product/common"
	"go-product/repositories"
	"go-product/services"
)

func main() {
	// 1. create iris instance
	app := iris.New()

	// 2. set error mode, prompt error in mvc mode
	app.Logger().SetLevel("debug")

	// 3. register template
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)

	// 4. set static template directory, first argument is route which is used with domain name
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))

	// Jump to the specified page when an exception occurs
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "Fail to load the page!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// Connect to database
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5. register controller, implement route
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	// register service into controller
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	orderRepository := repositories.NewOrderMangerRepository("order", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))

	// 6. initiate the service
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

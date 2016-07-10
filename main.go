package main

import "github.com/kataras/iris"

type UserAPI struct {
	*iris.Context
}

func main() {
	// api := iris.New()
	iris.Config.Render.Template.Engine = iris.PongoEngine
	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.Render("hi.html", map[string]interface{}{"Name": "natali"})
}

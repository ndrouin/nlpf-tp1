package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
)

func main() {
  iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
  iris.Get("/", newUser)
  iris.Listen(":80")
}

func newUser(ctx *iris.Context) {
  ctx.Render("new_user.html", nil)
}

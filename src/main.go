package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
)

func main() {


  iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
  iris.Get("/", accueil)
  iris.Get("/newUser", newUser)
  iris.Listen(":80")
}


func accueil(ctx *iris.Context) {

  ctx.MustRender("accueil.html", struct{ Name string }{Name: "iris1"}) 
  ctx.MustRender("accueil.html", struct{ Name string }{Name: "iris2"})
}

func newUser(ctx *iris.Context) {
  ctx.Render("new_user.html", nil)
}

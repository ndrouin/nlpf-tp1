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

ctx.Render("accueil_project_display.html", struct { Project string }{ Project: "Projet 1"}) 
  
ctx.Render("accueil_project_display.html", struct { Project string }{ Project: "Projet 2" }, iris.RenderOptions{"layout": iris.NoLayout})
}

func newUser(ctx *iris.Context) {
  ctx.Render("new_user.html", nil)
}

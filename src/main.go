package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
  _"github.com/go-sql-driver/mysql"
  "./model"
  "fmt"
)


func main() {
  model.InitModel()
  iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
  iris.Get("/", accueil)
  iris.Get("/newUser", newUser)
  iris.Get("/connection", connection)
  iris.Post("/registration", registration)
  iris.Post("/connection", auth)
  iris.Listen(":80")
}


//Fonction d'affichage de l'écran d'accueil et des projets
func accueil(ctx *iris.Context) {
  //la liste de tous les projets 

c := []string{"TEST3", "14/07/2017", "Ceci est la descriptionn de mon projet3", "3 euros"}

ctx.Render("Accueil/accueil_projects.html", map[string]interface{}{"Projects": c[0], "Dates": c[1], "Descriptions": c[2],
"Money":c[3]})}

func newUser(ctx *iris.Context) {
  ctx.Render("newUser.html", nil)
}

func connection(ctx *iris.Context) {
  ctx.Render("connexion.html", nil)
}



func registration(ctx *iris.Context) {
  //Get variables from form
  email := ctx.FormValueString("email")
  password := ctx.FormValueString("password")
  name := ctx.FormValueString("name")
  surname := ctx.FormValueString("surname")
  //call registration function from model
  model.Registration(email, password, name, surname)

  //return home page
  ctx.Render("connexion.html", nil)
}


func auth(ctx *iris.Context) {
  //Get variables from form
  email := ctx.FormValueString("email")
  password := ctx.FormValueString("password")
  //get result of authentification
  result := model.Connection(email, password)
  if result == true {
    ctx.Render("Accueil/accueil_title.html", nil)
  } else {
    ctx.Render("connexion_error.html", nil)
  }
}



















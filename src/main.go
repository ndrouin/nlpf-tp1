package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
  _"github.com/go-sql-driver/mysql"
  _"fmt"
  "./model"
)


func main() {
  model.InitModel()
  iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
  iris.Get("/", accueil)
  iris.Get("/newUser", newUser)
  iris.Post("/registration", registration)
  iris.Listen(":80")
}


//Fonction d'affichage de l'écran d'accueil et des projets
func accueil(ctx *iris.Context) {
  //la liste de tous les projets 
projects := []string{"Project 1", "Project 2"}

//Affichage du tite de la page d'accueil"
ctx.Render("Accueil/accueil_title.html", nil)

ctx.Render("Accueil/accueil_projects.html", map[string]interface{}{"Projects1":projects[0],  "Projects2":projects[1]})
//for _, c := range projects {

  //ctx.Render("Accueil/accueil_projects.html", struct{ Projects string }{Projects: c})




//ctx.Render("Accueil/accueil_projects.html", struct{ Projects string }{Projects: "Projet 1"}) 

//ctx.Render("Accueil/accueil_projects.html", struct{ Projects string }{Projects: "Projet 2"}, iris.RenderOptions{"layout": iris.NoLayout})
}

func newUser(ctx *iris.Context) {
  ctx.Render("newUser.html", nil)
}

func registration(ctx *iris.Context) {
  //Get variables from form
  email := ctx.FormValueString("email")
  password := ctx.FormValueString("password")
  //call registration function from model
  model.Registration(email, password)
}






















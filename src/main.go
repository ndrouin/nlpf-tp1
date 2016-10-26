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
  iris.Post("/registration", registration)
  iris.Listen(":80")
}


//Fonction d'affichage de l'écran d'accueil et des projets
func accueil(ctx *iris.Context) {
  //la liste de tous les projets 

project1 := []string{"TEST1", "12/05/1994", "Ceci est la descriptionn de mon projet1", " 1euros"}
project2 := []string{"TEST2", "13/06/2016", "Ceci est la descriptionn de mon projet2", "2 euros"}
project3 := []string{"TEST3", "14/07/2017", "Ceci est la descriptionn de mon projet3", "3 euros"}
project4 := []string{"TEST4", "15/08/2018", "Ceci est la descriptionn de mon projet4", "4 euros"}
project5 := []string{"TEST5", "16/09/2019", "Ceci est la descriptionn de mon projet5", "5 euros"}


projects := [][]string{project1, project2, project3, project4, project5}
fmt.Println(projects)
//Affichage du tite de la page d'accueil"
ctx.Render("Accueil/accueil_title.html", nil)

fmt.Println("c tipart")

for _,c:= range projects {
fmt.Println(projects)
ctx.Render("Accueil/accueil_projects.html", map[string]interface{}{"Projects": c[0], "Dates": c[1], "Descriptions": c[2],
"Money":c[3]}, iris.RenderOptions{"layout": iris.NoLayout})}
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






















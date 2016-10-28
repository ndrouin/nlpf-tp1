package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
  _"github.com/go-sql-driver/mysql"
  "./model"
  _"fmt"
  "strconv"
)

func main() {
  model.InitModel()
  iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
  iris.Get("/", home)
  iris.Get("/newUser", newUser)
  iris.Get("/connection", connection)
  iris.Post("/registration", registration)
  iris.Post("/connection", auth)
  my := iris.Party("/connect").Layout("layouts/layout_connected.html")
  {
    my.Get("/", home)
    my.Get("/newProject", newProject)
    my.Post("/addProject", addProject)
    my.Post("/addCounterpart", addCounterpart)
  }
  iris.Listen(":80")
}

//Display of the home page with all of the projects
func home(ctx *iris.Context) {
  model.DelOrphanCounterparts()
  projects := model.GetProjects()
  ctx.Render("home.html", struct { Projects []*model.Project}{Projects: projects})
}
//When the user wants to subscribe
func newUser(ctx *iris.Context) {
  ctx.Render("newUser.html", nil)
}

func newProject(ctx *iris.Context) {
  ctx.Render("newProject.html", nil)
}

func connection(ctx *iris.Context) {
  ctx.Render("connection.html", nil)
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
  ctx.Redirect("/")
}

func addProject(ctx *iris.Context) {
  //Get variables from form
  name := ctx.FormValueString("project_name")
  description := ctx.FormValueString("description")
  author := ctx.FormValueString("author_name")
  contact := ctx.FormValueString("email")
  //call AddProject function from model
  model.AddProject(name, description, author, contact)
  ctx.MustRender("notification.html", struct{ Text string }{Text: "Nouveau projet cree avec succes"})
}

func addCounterpart(ctx *iris.Context) {
  name := ctx.FormValueString("title")
  value, err := strconv.ParseInt(ctx.FormValueString("value"), 10, 64)
  description := ctx.FormValueString("description")
  model.AddCounterpart(name, value, description)
  _ = err
  type Vars struct {
    Add                   bool
    HasOrphanCounterpart  bool
    Counterparts          []*model.Counterpart
  }
  vars := Vars{
    Add:                  true,
    HasOrphanCounterpart: model.HasOrphanCounterpart(),
    Counterparts:         model.GetCounterparts(),
  }
  ctx.Render("newProject.html", vars)
}

func auth(ctx *iris.Context) {
  //Get variables from form
  email := ctx.FormValueString("email")
  password := ctx.FormValueString("password")
  //get result of authentification
  result := model.Connection(email, password)
  if result {
    ctx.Redirect("/connect/")
  } else {
    ctx.Render("connection_error.html", nil)
  }
}

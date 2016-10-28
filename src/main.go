package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
  _"github.com/go-sql-driver/mysql"
  "./model"
  _"fmt"
  "strconv"
  "strings"
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
    my.Get("/", homeConnect)
    my.Get("/newProject", newProject)
    my.Post("/addProject", addProject)
    my.Post("/addCounterpart", addCounterpart)
    my.Post("/participation", participation)
    my.Post("/selection", confirmation)
    my.Post("/gift", addSelection)
  }
  iris.Listen(":80")
}

func interfaceHome(add bool, connect bool) interface{} {
  model.DelOrphanCounterparts()
  projects := model.GetProjects()
  type Vars struct {
    Projects  []*model.Project
    Add       bool
    Connect   bool
  }
  vars := Vars {
    Projects: projects,
    Add:      add,
    Connect:   connect,
  }
  return vars
}
//Display of the home page with all of the projects
func home(ctx *iris.Context) {
  ctx.Render("home.html", interfaceHome(false, false))
}

func homeConnect(ctx *iris.Context) {
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
    ctx.Render("home.html", interfaceHome(false, true))
  }
}
//When the user wants to subscribe
func newUser(ctx *iris.Context) {
  ctx.Render("newUser.html", nil)
}

func newProject(ctx *iris.Context) {
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
    ctx.Render("newProject.html", nil)
  }
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
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
    //Get variables from form
    name := ctx.FormValueString("project_name")
    description := ctx.FormValueString("description")
    author := ctx.FormValueString("author_name")
    contact := ctx.FormValueString("email")
    //call AddProject function from model
    model.AddProject(name, description, author, contact)
    model.SetProjectCounterparts()
    ctx.Render("home.html", interfaceHome(true, true))
  }
}

func addCounterpart(ctx *iris.Context) {
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
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
}

func auth(ctx *iris.Context) {
  //Get variables from form
  email := ctx.FormValueString("email")
  password := ctx.FormValueString("password")
  //get result of authentification
  result := model.Connection(email, password)
  if result {
    ctx.Session().Set("isConnected", "true")
    ctx.Redirect("/connect/")
  } else {
    ctx.Render("connection_error.html", nil)
  }
}

func varsProject(id int64) interface{} {
  type Vars struct {
    Project       []*model.Project
    Counterparts  []*model.Counterpart
  }
  vars := Vars {
    Project:       model.GetProject(id),
    Counterparts:  model.GetProjectCounterparts(id),
  }
  return vars
}

func participation(ctx *iris.Context) {
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
    id := ctx.FormValueString("participation")
    i, err := strconv.ParseInt(id, 10, 64)
    _ = err
    ctx.Render("participation.html", varsProject(i))
  }
}
func confirmation(ctx *iris.Context) {
  id := ctx.FormValueString("selection")
  i, err := strconv.ParseInt(id, 10, 64)
  _ = err
  type Vars struct {
    Project       []*model.Project
    Counterpart  []*model.Counterpart
  }
  vars := Vars {
    Project:      model.GetProject(model.GetCounterpart(i)[0].Project),
    Counterpart:  model.GetCounterpart(i),
  }

  ctx.Render("confirmation.html", vars)
}


func addSelection(ctx *iris.Context) {
  if strings.Compare(ctx.Session().GetString("isConnected"), "true") == 0  {
    id := ctx.FormValueString("selection")
    i, err := strconv.ParseInt(id, 10, 64)
    _ = err
    model.AddSelection(i)
    ctx.Render("participation.html", varsProject(model.GetCounterpart(i)[0].Project))
  }
}














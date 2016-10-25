package main

import "github.com/kataras/iris"

func main() {

  iris.Get("/", newUser)
  iris.Listen(":80")
}

func newUser(ctx *iris.Context) {
  ctx.MustRender("new_user.html", anyBindingStruct{})
}

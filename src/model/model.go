package model

import (
 _ "github.com/go-sql-driver/mysql"
 "github.com/go-xorm/xorm"
)

type User struct {
      Email     string
      Password  string
}

var engine *xorm.Engine
var err error

func InitModel() {
  engine, err = xorm.NewEngine("mysql", "root:@/easywebsite")
  engine.Sync(new(User))
  _ = err
}

//add a new user in the DB
func Registration(email string, password string) {
  user := new(User)
  user.Email = email
  user.Password = password
  engine.Insert(user)
}

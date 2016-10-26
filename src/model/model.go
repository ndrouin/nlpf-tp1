package model

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
)

type User struct {
      Surname   string
      Name      string
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


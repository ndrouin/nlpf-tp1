package model

import (
 _ "github.com/go-sql-driver/mysql"
 "github.com/go-xorm/xorm"
)

type User struct {
      Email     string
      Password  string
}

func InitModel() {
  var err error
  var engine *xorm.Engine
  engine, err = xorm.NewEngine("mysql", "root:@/easywebsite")
  err = engine.Sync(new(User))
  _ = err
}

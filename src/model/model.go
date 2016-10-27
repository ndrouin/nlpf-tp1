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

type Project struct {
  Name        string
  Description string `xorm:"text"`
  Author      string
  Contact     string
}


var engine *xorm.Engine
var err error

func InitModel() {
  engine, err = xorm.NewEngine("mysql", "root:@/easywebsite")
  engine.Sync(new(User))
  engine.Sync(new(Project))
  _ = err
}


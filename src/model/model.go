package model

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
)


type User struct {
  Id        int64 `xorm:"id pk not null autoincr"`
  Surname   string
  Name      string
  Email     string `xorm: "unique"`
  Password  string
}

type Project struct {
  Id          int64 `xorm:"id pk not null autoincr"`
  Name        string
  Description string `xorm:"text"`
  Author      string
  Contact     string
  Price       int64
  Creation    string
}

type Counterpart struct {
  Id          int64 `xorm:"id pk not null autoincr"`
  Name        string
  Value       int64
  Description string `xorm:"text"`
  Project     int64
}


var engine *xorm.Engine
var err error

func InitModel() {
  engine, err = xorm.NewEngine("mysql", "root:@/easywebsite")
  engine.Sync(new(User))
  engine.Sync(new(Project))
  engine.Sync(new(Counterpart))
  _ = err
}

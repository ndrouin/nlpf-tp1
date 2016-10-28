package model

import (
  "crypto/sha256"
  "encoding/hex"
  _"fmt"
)
//add a new user in the DB
func Registration(email string, password string, name string, surname string) {
  //initiate new user
  user := new(User)
  user.Email = email
  user.Password = crypt(password)
  user.Name = name
  user.Surname = surname
  //insert new project in DB
  engine.Insert(user)
}

func Connection(email string, password string) bool {
  user := &User{
    Email: email,
    Password: crypt(password),
  }
  has, err := engine.Get(user)
  _ = err
  return has
}

func crypt(str string) string {
  h:= sha256.Sum256([]byte(str))
  return hex.EncodeToString(h[:])
}

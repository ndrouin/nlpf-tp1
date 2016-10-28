package model

import (
  "time"
  "strconv"
)

func AddProject(name string, description string, author string, contact string) {
  //initiate new project
  project := new(Project)
  project.Name = name
  project.Description = description
  project.Author = author
  project.Contact = contact
  project.Price = 0
  year, month, day := time.Now().Date()
  date := strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)
  project.Creation = date
  //insert new project in DB
  engine.Insert(project)
}

func AddCounterpart(name string, value int64, description string) {
  counterpart := new(Counterpart)
  counterpart.Name = name
  counterpart.Value = value
  counterpart.Description = description
  engine.Insert(counterpart)
}


func GetProjects() []*Project {
  var projects []*Project
  engine.Find(&projects)
  return projects
}

func HasOrphanCounterpart() bool {
  counterpart := &Counterpart{
    Project: 0,
  }
  has, err := engine.Get(counterpart)
  _ = err
  return has
}










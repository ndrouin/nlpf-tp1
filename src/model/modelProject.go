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
  engine.Desc("id").Find(&projects)
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

func GetCounterparts() []*Counterpart {
  var counterparts []*Counterpart
  engine.Having("project=0").Find(&counterparts)
  return counterparts
}

func DelOrphanCounterparts() {
  engine.Query("DELETE FROM counterpart WHERE project=0")
}

func SetProjectCounterparts() {
  var projects []*Project
  engine.Desc("id").Find(&projects)
  _ = err
  index := projects[0].Id
  sql := "UPDATE counterpart SET project=" + strconv.FormatInt(index, 10) + " WHERE project=0"
  engine.Query(sql)
}

func GetProject(id int64) []*Project{
  var project []*Project
  engine.Having("id="+strconv.FormatInt(id, 10)).Find(&project)
  return project
}

func GetProjectCounterparts(id int64) []*Counterpart{
  var counterparts []*Counterpart
  engine.Having("project="+strconv.FormatInt(id, 10)).Find(&counterparts)
  return counterparts
}

func AddSelection(id int64) {
  var counterpart []*Counterpart
  engine.Having("id="+strconv.FormatInt(id, 10)).Find(&counterpart)
  value := counterpart[0].Value
  var project []*Project
  engine.Having("id="+strconv.FormatInt(counterpart[0].Project, 10)).Find(&project)
  price := project[0].Price + value
  sql := "UPDATE project SET price=" + strconv.FormatInt(price, 10) + " WHERE id=" +strconv.FormatInt(counterpart[0].Project, 10)
  engine.Query(sql)

}

func GetCounterpart(id int64) []*Counterpart{
  var counterpart []*Counterpart
  engine.Having("id="+strconv.FormatInt(id, 10)).Find(&counterpart)
  return counterpart
}

func BestProjects() []*Project {
  var projects []*Project
  engine.Desc("price").Find(&projects)
  return projects
}



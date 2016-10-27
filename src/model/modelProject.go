package model

func AddProject(name string, description string, author string, contact string) {
  //initiate new project
  project := new(Project)
  project.Name = name
  project.Description = description
  project.Author = author
  project.Contact = contact
  //insert new project in DB
  engine.Insert(project)
}

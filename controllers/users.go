package controllers

import (
  "github.com/altonymous/forum/models"
  "github.com/hoisie/web"
  "log"
)

type Users struct{}

func (self *Users) Index(webContext *web.Context) {
  log.Println("user.Index")
  // generateResourceRoutes(Controllers(&controllers.Forums{}))
  users := models.User{}.All()

  renderTemplate(webContext, "templates/users", "index", &users, 200)
}

func (self *Users) Show(webContext *web.Context) {
  log.Println("user.Show")
}

func (self *Users) New(webContext *web.Context) {
  log.Println("user.New")
  renderTemplate(webContext, "templates/users", "new", nil, 200)
}

func (self *Users) Edit(webContext *web.Context) {
  log.Println("user.Edit")
}

func (self *Users) Create(webContext *web.Context) {
  log.Println("user.Create")

  user := models.User{}.Create(webContext.Params)

  if user.Errors != nil {
    renderTemplate(webContext, "templates/users", "new", &user, 400)
  }

  users := models.User{}.All()
  renderTemplate(webContext, "templates/users", "index", &users, 201)
}

func (self *Users) Update(webContext *web.Context) {
  log.Println("user.Update")
}

func (self *Users) Delete(webContext *web.Context) {
  log.Println("user.Delete")
}

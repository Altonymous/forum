package controllers

import (
  "github.com/altonymous/forum/models"
  "github.com/hoisie/web"
  "log"
)

type Forums struct{}

func (self *Forums) Index(webContext *web.Context) {
  log.Println("forum.Index")
  // generateResourceRoutes(Controllers(&controllers.Forums{}))
  forums := models.Forum{}.All()

  renderTemplate(webContext, "templates/forums", "index", &forums, 200)
}

func (self *Forums) Show(webContext *web.Context) {
  log.Println("forum.Show")
}

func (self *Forums) New(webContext *web.Context) {
  log.Println("forum.New")
  renderTemplate(webContext, "templates/forums", "new", nil, 200)
}

func (self *Forums) Edit(webContext *web.Context) {
  log.Println("forum.Edit")
}

func (self *Forums) Create(webContext *web.Context) {
  log.Println("forum.Create")

  forum := models.Forum{}.Create(webContext.Params)

  if forum.Errors != nil {
    renderTemplate(webContext, "templates/forums", "new", &forum, 400)
  }

  forums := models.Forum{}.All()
  renderTemplate(webContext, "templates/forums", "index", &forums, 201)
}

func (self *Forums) Update(webContext *web.Context) {
  log.Println("forum.Update")
}

func (self *Forums) Delete(webContext *web.Context) {
  log.Println("forum.delete")
}

package controllers

import (
  "github.com/altonymous/forum/models"
  "github.com/hoisie/web"
  "log"
)

type Posts struct{}

func (self *Posts) Index(webContext *web.Context) {
  log.Println("post.Index")
  // generateResourceRoutes(Controllers(&controllers.Forums{}))
  posts := models.Post{}.All()

  renderTemplate(webContext, "templates/posts", "index", &posts, 200)
}

func (self *Posts) Show(webContext *web.Context) {
  log.Println("post.Show")
}

func (self *Posts) New(webContext *web.Context) {
  log.Println("post.New")
  renderTemplate(webContext, "templates/posts", "new", nil, 200)
}

func (self *Posts) Edit(webContext *web.Context) {
  log.Println("post.Edit")
}

func (self *Posts) Create(webContext *web.Context) {
  log.Println("post.Create")

  post := models.Post{}.Create(webContext.Params)

  if post.Errors != nil {
    renderTemplate(webContext, "templates/posts", "new", &post, 400)
  }

  posts := models.Post{}.All()
  renderTemplate(webContext, "templates/posts", "index", &posts, 201)
}

func (self *Posts) Update(webContext *web.Context) {
  log.Println("post.Update")
}

func (self *Posts) Delete(webContext *web.Context) {
  log.Println("post.Delete")
}

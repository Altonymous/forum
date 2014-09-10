package controllers

import (
	"github.com/altonymous/forum/models"
	"github.com/hoisie/web"
	"log"
)

type Topics struct{}

func (self *Topics) Index(webContext *web.Context) {
	log.Println("topic.Index")
	topics := models.Topic{}.All()

	renderTemplate(webContext, "templates/topics", "index", &topics, 200)
}

func (self *Topics) Show(webContext *web.Context) {
	log.Println("topic.Show")
}

func (self *Topics) New(webContext *web.Context) {
	log.Println("topic.New")
	renderTemplate(webContext, "templates/topics", "new", nil, 200)
}

func (self *Topics) Edit(webContext *web.Context) {
	log.Println("topic.Edit")
}

func (self *Topics) Create(webContext *web.Context) {
	log.Println("topic.Create")

	topic := models.Topic{}.Create(webContext.Params)

	if topic.Errors != nil {
		renderTemplate(webContext, "templates/topics", "new", &topic, 400)
	}

	topics := models.Topic{}.All()
	renderTemplate(webContext, "templates/topics", "index", &topics, 201)
}

func (self *Topics) Update(webContext *web.Context) {
	log.Println("topic.Update")
}

func (self *Topics) Delete(webContext *web.Context) {
	log.Println("topic.Delete")
}

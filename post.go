package main

import (
  "github.com/hoisie/web"
  "time"
)

type post struct {
  Identity  int
  Message   string
  Poster    user
  CreatedAt time.Time
}

func (self *post) index(webContext *web.Context) {

}

func (self *post) show(webContext *web.Context) {

}

func (self *post) new(webContext *web.Context) {

}

func (self *post) edit(webContext *web.Context) {

}

func (self *post) create(webContext *web.Context) {

}

func (self *post) update(webContext *web.Context) {

}

func (self *post) delete(webContext *web.Context) {

}

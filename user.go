package main

import (
  "github.com/hoisie/web"
  "time"
)

type user struct {
  Identity  int
  Username  string
  PassHash  string
  FirstName string
  LastName  string
  Email     string
  Handle    string
  CreatedAt time.Time
}

func (self *user) index(webContext *web.Context) {

}

func (self *user) create(webContext *web.Context) {

}

func (self *user) show(webContext *web.Context) {

}

func (self *user) update(webContext *web.Context) {

}

func (self *user) delete(webContext *web.Context) {

}

package main

import (
  "github.com/hoisie/web"
  "time"
)

type topic struct {
  Identity     int
  Name         string
  LatestPost   post
  LatestPoster user
  CreatedAt    time.Time
}

func (self *topic) index(webContext *web.Context) {

}

func (self *topic) create(webContext *web.Context) {

}

func (self *topic) show(webContext *web.Context) {

}

func (self *topic) update(webContext *web.Context) {

}

func (self *topic) delete(webContext *web.Context) {

}

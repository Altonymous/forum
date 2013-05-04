package main

import (
  "github.com/hoisie/web"
  "time"
)

type forum struct {
  Identity    int
  Name        string
  Description string
  LatestTopic topic
  CreatedAt   time.Time
}

func (self *forum) index(webContext *web.Context) {

}

func (self *forum) create(webContext *web.Context) {

}

func (self *forum) show(webContext *web.Context) {

}

func (self *forum) update(webContext *web.Context) {

}

func (self *forum) delete(webContext *web.Context) {

}

package main

import (
  "bitbucket.org/pkg/inflect"
  "github.com/hoisie/web"
  "log"
  "reflect"
  "strings"
)

type resource interface {
  index(webContext *web.Context)
  show(webContext *web.Context)
  create(webContext *web.Context)
  update(webContext *web.Context)
  delete(webContext *web.Context)
}

func generateRoutes() {
  log.Println("starting route generation")
  generateResourceRoutes(resource(&forum{}))
  generateResourceRoutes(resource(&topic{}))
  generateResourceRoutes(resource(&post{}))
  generateResourceRoutes(resource(&user{}))
  log.Println("finished route generation")
}

func generateResourceRoutes(resource resource) {
  resourceString := inflect.Pluralize(friendlyName(resource))

  if !globalConfiguration.Database.tableExists(resourceString) {
    globalConfiguration.Database.tableCreate(resourceString)
  }

  log.Printf(" - starting /%s routes generation\n", resourceString)
  web.Get("(?i)/"+resourceString+"/?", resource.index)
  web.Get("(?i)/"+resourceString+"/(\\d*)", resource.show)
  web.Post("(?i)/"+resourceString+"/?", resource.create)
  web.Put("(?i)/"+resourceString+"/(\\d*)", resource.update)
  web.Delete("(?i)/"+resourceString+"/(\\d*)", resource.delete)
  log.Printf(" - finished /%s routes generation\n", resourceString)
}

func startServer(port string) {
  generateRoutes()

  web.Run("0.0.0.0:" + port)
}

func friendlyName(resource resource) string {
  return strings.Split(reflect.TypeOf(resource).String(), ".")[1]
}

package main

import (
  "github.com/altonymous/forum/controllers"
  "github.com/altonymous/forum/models"
  "github.com/hoisie/web"
  "log"
)

func generateRoutes() {
  log.Println("starting route generation")
  generateResourceRoutes(&controllers.Forums{})
  generateResourceRoutes(&controllers.Topics{})
  generateResourceRoutes(&controllers.Posts{})
  generateResourceRoutes(&controllers.Users{})
  log.Println("finished route generation")
}

func generateResourceRoutes(resource controllers.Controllers) {
  resourceString := controllers.FriendlyName(resource)

  models.SetupTable(resourceString)

  log.Printf(" - starting /%s routes generation\n", resourceString)
  web.Get("(?i)/"+resourceString+"/?", resource.Index)
  web.Get("(?i)/"+resourceString+"/(\\d*)", resource.Show)
  web.Get("(?i)/"+resourceString+"/new/?", resource.New)
  web.Get("(?i)/"+resourceString+"/(\\d*)", resource.Edit)

  web.Post("(?i)/"+resourceString+"/?", resource.Create)
  web.Put("(?i)/"+resourceString+"/(\\d*)", resource.Update)
  web.Delete("(?i)/"+resourceString+"/(\\d*)", resource.Delete)
  log.Printf(" - finished /%s routes generation\n", resourceString)
}

func startServer(port string) {
  models.SetDatabase(&globalConfiguration.Database)

  web.Get("/?", index)
  generateRoutes()

  web.Run("0.0.0.0:" + port)
}

func index(webContext *web.Context) {
  webContext.WriteString("Powered by <a href=\"http://www.github.com/Altonymous/forums/\">Altonymous forums!</a>")
}

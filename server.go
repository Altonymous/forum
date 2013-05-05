package main

import (
  "bitbucket.org/pkg/inflect"
  "github.com/hoisie/web"
  "html/template"
  "io/ioutil"
  "log"
  "path/filepath"
  "reflect"
  "strings"
)

type resource interface {
  index(webContext *web.Context)
  show(webContext *web.Context)
  new(webContext *web.Context)
  edit(webContext *web.Context)

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
  web.Get("(?i)/"+resourceString+"/new/?", resource.new)
  web.Get("(?i)/"+resourceString+"/(\\d*)", resource.edit)

  web.Post("(?i)/"+resourceString+"/?", resource.create)
  web.Put("(?i)/"+resourceString+"/(\\d*)", resource.update)
  web.Delete("(?i)/"+resourceString+"/(\\d*)", resource.delete)
  log.Printf(" - finished /%s routes generation\n", resourceString)
}

func startServer(port string) {
  web.Get("/?", index)
  generateRoutes()
  cacheTemplates("templates")

  web.Run("0.0.0.0:" + port)
}

var templates = make(map[string]*template.Template)

func cacheTemplates(templateFilePath string) {
  fileInfos, err := ioutil.ReadDir(templateFilePath)

  if err != nil {
    log.Println("[ERROR] cacheTemplates: ", err)
    panic(err)
  }

  for i := 0; i < len(fileInfos); i++ {
    if fileInfos[i].IsDir() {
      cacheTemplates(templateFilePath + "/" + fileInfos[i].Name())
    } else {
      if filepath.Ext(fileInfos[i].Name()) == ".html" {
        if templates[templateFilePath] == nil {
          templates[templateFilePath] = template.New(templateFilePath)
        }

        _, err := templates[templateFilePath].ParseFiles(templateFilePath + "/" + fileInfos[i].Name())

        if err != nil {
          log.Println("[ERROR] cacheTemplates: ", err)
          panic(err)
        }
      }
    }
  }
}

func renderTemplate(webContext *web.Context, template string, action string, context interface{}, statusCode int) {
  webContext.WriteHeader(statusCode)
  err := templates[template].ExecuteTemplate(webContext, action+".html", context)

  if err != nil {
    log.Println("[ERROR] renderTemplate: ", err)
    webContext.Abort(500, "Unable to process request.")
  }
}

func friendlyName(resource resource) string {
  return strings.Split(reflect.TypeOf(resource).String(), ".")[1]
}

func index(webContext *web.Context) {
  webContext.WriteString("Powered by <a href=\"http://www.github.com/Altonymous/forums/\">Altonymous forums!</a>")
}

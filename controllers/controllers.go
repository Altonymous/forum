package controllers

import (
	"github.com/hoisie/web"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strings"
)

type Controllers interface {
	Index(webContext *web.Context)
	Show(webContext *web.Context)
	New(webContext *web.Context)
	Edit(webContext *web.Context)

	Create(webContext *web.Context)
	Update(webContext *web.Context)
	Delete(webContext *web.Context)
}

var templates = make(map[string]*template.Template)

func init() {
	cacheTemplates("templates")
}

func renderTemplate(webContext *web.Context, template string, action string, context interface{}, statusCode int) {
	webContext.WriteHeader(statusCode)
	err := templates[template].ExecuteTemplate(webContext, action+".html", context)

	if err != nil {
		log.Println("[ERROR] renderTemplate: ", err)
		webContext.Abort(500, "Unable to process request.")
	}
}

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

func FriendlyName(controller Controllers) string {
	return strings.ToLower(strings.Split(reflect.TypeOf(controller).String(), ".")[1])
}

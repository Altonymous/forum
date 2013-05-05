package main

import (
  "crypto/sha512"
  "fmt"
  "github.com/christopherhesse/rethinkgo"
  "github.com/hoisie/web"
  "io"
  "log"
  "time"
)

type user struct {
  Id           string    `json:"id"`
  Username     string    `json:"username"`
  PasswordHash string    `json:"password_hash"`
  FirstName    string    `json:"first_name"`
  LastName     string    `json:"last_name"`
  Email        string    `json:"email"`
  Handle       string    `json:"handle"`
  CreatedAt    time.Time `json:"created_at`
  Errors       map[string]string
}

func (self *user) all() *[]user {
  log.Println("user.all")
  session := globalConfiguration.Database.getSession()

  var users []user
  err := rethinkgo.Table("users").Run(session).All(&users)

  if err != nil {
    log.Println("[ERROR] user.index: ", err)
    panic(err)
  }

  return &users
}

func (self *user) findById(id int) *user {
  return self
}

func (self *user) index(webContext *web.Context) {
  log.Println("user.index")
  users := self.all()

  renderTemplate(webContext, "templates/users", "index", &users, 200)
}

func (self *user) show(webContext *web.Context) {
  log.Println("user.show")
}

func (self *user) new(webContext *web.Context) {
  log.Println("user.new")
  renderTemplate(webContext, "templates/users", "new", nil, 200)
}

func (self *user) edit(webContext *web.Context) {
  log.Println("user.edit")
}

func (self *user) create(webContext *web.Context) {
  log.Println("user.create")
  // Setup session
  session := globalConfiguration.Database.getSession()

  // Setup table
  table := rethinkgo.Table("users")

  params := webContext.Params

  var myUser user
  if params["password"] != params["password_confirmation"] {
    errors := make(map[string]string)
    errors["password"] = "do not match"

    myUser = user{Username: params["username"],
        FirstName: params["first_name"],
        LastName:  params["last_name"],
        Email:     params["email"],
        Handle:    params["handle"],
        Errors:    errors}

    renderTemplate(webContext, "templates/users", "new", &myUser, 400)
    return
  }

  passwordHash := sha512.New()
  io.WriteString(passwordHash, params["password"])

  user := rethinkgo.Map{
    "username":      params["username"],
    "password_hash": fmt.Sprintf("%x", passwordHash.Sum(nil)),
    "first_name":    params["first_name"],
    "last_name":     params["last_name"],
    "email":         params["email"],
    "handle":        params["handle"],
    "created_at":    time.Now()}

  // insert the reading
  var response rethinkgo.WriteResponse
  err := table.Insert(user).Run(session).One(&response)

  if err != nil {
    log.Println("[ERROR] user.create: ", err)
    panic(err)
  }

  renderTemplate(webContext, "templates/users", "index", self.all(), 201)
}

func (self *user) update(webContext *web.Context) {
  log.Println("user.update")
}

func (self *user) delete(webContext *web.Context) {
  log.Println("user.delet")
}

func (self *user) Name() string {
  return self.FirstName + " " + self.LastName
}

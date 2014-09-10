package models

import (
	"github.com/christopherhesse/rethinkgo"
	"log"
	"time"
)

type User struct {
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

func (self User) All() []User {
	log.Println("User.All")
	session := database.getSession()

	var users []User
	err := rethinkgo.Table("users").Run(session).All(&users)

	if err != nil {
		log.Println("[ERROR] User.All: ", err)
		panic(err)
	}

	return users
}

func (self User) FindById(id string) User {
	log.Println("User.FindById")
	session := database.getSession()

	var user User
	err := rethinkgo.Table("users").Get(id).Run(session).One(&user)

	if err != nil {
		log.Println("[ERROR] User.All: ", err)
		panic(err)
	}

	return user
}

func (self User) Create(params map[string]string) User {
	// Setup session
	session := database.getSession()

	// Setup table
	table := rethinkgo.Table("users")

	errors := make(map[string]string)
	user := User{Username: params["username"],
		FirstName: params["first_name"],
		LastName:  params["last_name"],
		Email:     params["email"],
		Handle:    params["handle"]}

	if params["password"] != params["password_confirmation"] {
		errors["password"] = "do not match"

		user.Errors = errors
	} else {
		userMap := rethinkgo.Map{
			"username":      params["username"],
			"password_hash": passwordHash(params["password"]),
			"first_name":    params["first_name"],
			"last_name":     params["last_name"],
			"email":         params["email"],
			"handle":        params["handle"],
			"created_at":    time.Now()}

		// insert the reading
		var response rethinkgo.WriteResponse
		err := table.Insert(userMap).Run(session).One(&response)

		if err != nil {
			log.Println("[ERROR] User.Create: ", err)
			errors["unknown"] = err.Error()
			user.Errors = errors
		}
	}

	return user
}

func (self User) Name() string {
	return self.FirstName + " " + self.LastName
}

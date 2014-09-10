package models

import (
	"time"
)

type Post struct {
	Id        string            `json:"id"`
	Message   string            `json:"message"`
	Poster    User              `json:"poster"`
	CreatedAt time.Time         `json:"created_at"`
	Errors    map[string]string `json:"errors"`
}

func (self Post) All() []Post {
	return []Post{{Message: "Sexah"}}
}

func (self Post) FindById() Post {
	return Post{Message: "Sexah"}
}

func (self Post) Create(params map[string]string) Post {
	return Post{Message: "Sexah"}
}

package models

import (
	"time"
)

type Forum struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	LatestTopic Topic             `json:"latest_topic"`
	CreatedAt   time.Time         `json:"created_at"`
	Errors      map[string]string `json:"errors"`
}

func (self Forum) All() []Forum {
	return []Forum{{Name: "Sexah"}}
}

func (self Forum) FindById() Forum {
	return Forum{Name: "Sexah"}
}

func (self Forum) Create(params map[string]string) Forum {
	return Forum{Name: "Sexah"}
}

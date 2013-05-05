package models

import (
  "time"
)

type Topic struct {
  Id           string            `json:"id"`
  Name         string            `json:"name"`
  LatestPost   Post              `json:"latest_post"`
  LatestPoster User              `json:"latest_poster"`
  CreatedAt    time.Time         `json:"created_at"`
  Errors       map[string]string `json:"errors"`
}

func (self Topic) All() []Topic {
  return []Topic{{Name: "Sexah"}}
}

func (self Topic) FindById() Topic {
  return Topic{Name: "Sexah"}
}

func (self Topic) Create(params map[string]string) Topic {
  return Topic{Name: "Sexah"}
}

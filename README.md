forum
=====

Web Forum built in go (#golang)

Installation
=====
go get -u github.com/christopherhesse/rethinkgo
go get -u github.com/hoisie/mustache
go get -u github.com/hoisie/web
go get -u bitbucket.org/pkg/inflect

Structure
=====
User
  has_many Topics

Forum
  has_many Topics

Topic
  belongs_to User
  belongs_to Forum

  has_many Posts

Post
  belongs_to User
  belongs_to Topic
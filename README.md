forum.go
-----

Basic Web Forum built with go (#golang)

Installation
------------

### Prerequisites:

    brew install mercurial # If you do not have [mercurial] installed, it is required:

### Dependencies

    go get -u github.com/hoisie/web
    go get -u github.com/christopherhesse/rethinkgo


Model Structure
-----
* User
 * has_many Topics
* Forum
 * has_many Topics
* Topic
 * belongs_to User
 * belongs_to Forum
 * has_many Posts
* Post
 * belongs_to User
 * belongs_to Topic

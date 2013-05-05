forum.go
--------

Basic Web Forum built with go (#golang)

Installation
------------

### Prerequisites:

    brew install mercurial
    brew install rethinkdb

### Dependencies

    go get -u github.com/hoisie/web
    go get -u github.com/christopherhesse/rethinkgo

Running
-------

    First start your rethinkdb by simply typing `rethinkdb` at a command prompt.

    Navigate to http://localhost:8080 in your browser and create a new database called `community`

    In order to run this application you'll need to update the conf file to point to your database, by default the configuration file is set to rethinkdb defaults.  If you change the defaults you will need to update the conf file.


Model Structure
---------------

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

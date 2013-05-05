package main

import (
  "github.com/christopherhesse/rethinkgo"
  "log"
)

type database struct {
  Host    string
  Port    string
  Name    string
  Session *rethinkgo.Session
}

func (database *database) getSession() *rethinkgo.Session {
  if database.Session == nil {
    var err error
    database.Session, err = rethinkgo.Connect(globalConfiguration.Database.Host+":"+globalConfiguration.Database.Port, globalConfiguration.Database.Name)

    if err != nil {
      log.Println("[ERROR] getSession: ", err)
      panic(err) // no, not really
    }
  }

  return database.Session
}

func (database *database) tableExists(tableName string) bool {
  log.Printf(" - checking if %s table exists\n", tableName)
  session := database.getSession()

  var tables []string
  err := rethinkgo.TableList().Run(session).One(&tables)

  if err != nil {
    log.Println("[ERROR] tableExists: ", err)
    panic(err)
  }

  for i := 0; i < len(tables); i++ {
    if tableName == tables[i] {
      return true
    }
  }

  return false
}

func (database *database) tableCreate(tableName string) {
  log.Printf(" - creating %s table\n", tableName)
  session := database.getSession()

  err := rethinkgo.TableCreate(tableName).Run(session).Exec()

  if err != nil {
    log.Println("[ERROR] tableCreate", err)
    panic(err)
  }
}

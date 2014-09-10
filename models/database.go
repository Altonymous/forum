package models

import (
	"github.com/christopherhesse/rethinkgo"
	"log"
)

type Database struct {
	Host    string
	Port    string
	Name    string
	Session *rethinkgo.Session
}

var database *Database = new(Database)

func SetDatabase(_database *Database) {
	database = _database
}

func (self *Database) getSession() *rethinkgo.Session {
	if self.Session == nil {
		var err error
		self.Session, err = rethinkgo.Connect(self.Host+":"+self.Port, self.Name)

		if err != nil {
			log.Println("[ERROR] getSession: ", err)
			panic(err) // no, not really
		}
	}

	return self.Session
}

func SetupTable(tableName string) {
	if !database.tableExists(tableName) {
		database.tableCreate(tableName)
	}
}

func (self *Database) tableExists(tableName string) bool {
	log.Printf(" - checking if %s table exists\n", tableName)
	session := self.getSession()

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

func (self *Database) tableCreate(tableName string) {
	log.Printf(" - creating %s table\n", tableName)
	session := self.getSession()

	err := rethinkgo.TableCreate(tableName).Run(session).Exec()

	if err != nil {
		log.Println("[ERROR] tableCreate", err)
		panic(err)
	}
}

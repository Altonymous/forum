package main

import (
    "fmt"
    r "github.com/christopherhesse/rethinkgo"
)

func main() {
    session, err := r.Connect("localhost:28015", "test")
    if err != nil {
        fmt.Println(err)
    }
    var tables []string
    err = r.TableList().Run(session).One(&tables)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(tables)
}
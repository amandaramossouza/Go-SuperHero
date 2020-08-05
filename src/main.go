package main

import (
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
    "controllers"
    "dao"
    //"github.com/gorilla/mux"
)

func Routes() {
    dao.InitDB()
    defer dao.CloseDB()

    http.HandleFunc("/", controllers.HomePage)
    http.HandleFunc("/all", controllers.AllPage)
    http.HandleFunc("/heroes", controllers.HeroesPage)
    http.HandleFunc("/villains", controllers.VillainsPage)

    fmt.Println("Api-rest - port 3003")
    log.Fatal(http.ListenAndServe(":3003", nil))
}

func main() {
    Routes()
}
package controllers

import (
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	
	//"strconv"
	"dao"
	//"models"
)


func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    //fmt.Println("Endpoint: homePage")
}

//
func AllPage(w http.ResponseWriter, r *http.Request){
	heroes, err := dao.GetAllHeroes()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	FormatResponseToJSON(w, http.StatusOK, heroes)
    //fmt.Fprintf(w, "Welcome to all Page!")
    fmt.Println("Endpoint: all")

}

func HeroesPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to heroes Page!")
    fmt.Println("Endpoint: heroes")
}

func VillainsPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to villains Page!")
    fmt.Println("Endpoint: villains")
}

//FormatResponseToJSON format the response to json
func FormatResponseToJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
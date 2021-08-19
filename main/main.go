package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Sudarshan",
			Title: "Student",
			State: "Puducherry",
		},
	}
	// encoding it to JSON
	json.NewEncoder(response).Encode(who)
	// printing to output
	fmt.Println("Endpoint Hit", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "こんにちわ to my custom Go Web API!!")
	fmt.Println("Endpoint Hit: HomePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Sudarshan"
	fmt.Fprintf(response, "A little bit about Sudarshan... :-) ")
	fmt.Println("Endpoint Hit", who)
}

func request() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)
	// starting the server at port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	models "goplayground/objects/models"
	service "goplayground/objects/services"
	"log"
	"net/http"
)

func FindPeopleOver(person models.Person) bool {
	return person.Age > 1
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {

	var thing = "hi there world"
	fmt.Println(thing)
	var container = service.Container[models.Person]{
		People: []models.Person{{Message: "hi", Age: 30}},
	}
	var peopleOver1 = container.Filter(FindPeopleOver)
	fmt.Println(peopleOver1)
	responseWriter.Write([]byte("HI there world"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body Demon
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	err = CreateDemon(body)
	var data Response
	if err != nil {
		data = Response{
			Success: false,
			Data:    Data{"Error"},
		}
	} else {
		data = Response{
			Success: true,
			Data:    Data{"Hello"},
		}
	}
	Respond(w, r, http.StatusOK, data)
}

func getDemon(w http.ResponseWriter, r *http.Request) {
	demon, err := FindOneDemon(bson.M{"data": "Vimal"})
	var data Response
	if err != nil {
		data = Response{
			Success: false,
			Data:    Data{"Error"},
		}
	} else {
		fmt.Println(demon)
		data = Response{
			Success: true,
			Data:    demon,
		}
	}
	Respond(w, r, http.StatusOK, data)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	data := Response{
		Success: true,
		Data:    Data{"Hello"},
	}
	Respond(w, r, http.StatusOK, data)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/create", create)
	http.HandleFunc("/get", getDemon)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("listeneing ")
		log.Fatal("ListenAndServe:- ", err)
	} else {
		fmt.Println("Error in listenings")
		log.Fatal("ERROR")
	}
}

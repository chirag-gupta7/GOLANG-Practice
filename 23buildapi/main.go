package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Cource -File
type Cource struct {
	CourceName  string  `json:"coursename"`
	CourceId    string  `json:"courseid"`
	CourcePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database
var cources []Cource

func (c *Cource) IsEmpty() bool {
	return c.CourceName == "" && c.CourceId == ""
}

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API</h1>"))

}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}

func getOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Cource")
	w.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(r)

	for _, cource := range cources {
		if cource.CourceId == parameters["id"] {
			json.NewEncoder(w).Encode(cource)
			return
		}
	}
	json.NewEncoder(w).Encode("No Cource found in given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Cource")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found in request body")
		return
	}

	var cource Cource

	_ = json.NewDecoder(r.Body).Decode(&cource)

	if cource.IsEmpty() {
		json.NewEncoder(w).Encode("No data found in request body")
		return
	}

	// rand.Seed(time.Now().UnixNano())

	cource.CourceId = strconv.Itoa(rand.IntN(100))
	cources = append(cources, cource)
	json.NewEncoder(w).Encode(cource)
	return

}

func updataOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Cource")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the request
	params := mux.Vars(r)

	//loop, id, remove, update
	for index, cource := range cources {
		if cource.CourceId == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			var cource Cource
			_ = json.NewDecoder(r.Body).Decode(&cource)
			cource.CourceId = params["id"]
			cources = append(cources, cource)
			json.NewEncoder(w).Encode(cources)
			return

		}
	}
	json.NewEncoder(w).encode("No Cource found in given ID")
	return

}

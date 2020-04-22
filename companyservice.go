package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/beevik/guid"
)

func main() {
	println("This is the company service.")
	println("It is made in Golang!")


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	http.ListenAndServe(":8090", router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createCompany(w http.ResponseWriter, r *http.Request)  {
	var newCompany company
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "There was an error")
	}
	json.Unmarshal(reqBody, &newCompany)

	newCompany.Id = *guid.New()

	companies = append(companies, newCompany)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCompany)
}



type company struct {
	Id   			guid.Guid `json:"Id"`
	Name 			string `json:"Name"`
}

type allCompanies []company

var companies = allCompanies{
	{
		Name: "Shirwood's Kitchen",
	},
}
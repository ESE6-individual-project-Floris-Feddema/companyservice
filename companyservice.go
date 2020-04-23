package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beevik/guid"
	"github.com/gorilla/mux"
)

func main() {
	println("This is the company service.")
	println("It is made in Golang!")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/companies", createCompany).Methods("POST")
	router.HandleFunc("/companies", getAllCompanies).Methods("GET")
	router.HandleFunc("/companies/{id}", getOneCompany).Methods("GET")
	router.HandleFunc("/companies/{id}", updateCompany).Methods("PUT", "PATCH")
	router.HandleFunc("/companies/{id}", deleteCompany).Methods("DELETE")
	http.ListenAndServe(":8090", router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createCompany(w http.ResponseWriter, r *http.Request) {
	var newCompany Company
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

func getOneCompany(w http.ResponseWriter, r *http.Request)  {
	companyId := mux.Vars(r)["id"]

	for _, singleCompany := range companies {
		if singleCompany.Id.String() == companyId {
			json.NewEncoder(w).Encode(singleCompany)
		}
	}
}

func getAllCompanies(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(companies)
}

func updateCompany(w http.ResponseWriter, r *http.Request){
	companyId := mux.Vars(r)["id"]
	var updatedCompany Company
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Yo dis not good")
	}
	json.Unmarshal(reqBody, &updatedCompany)

	for i, singleCompany := range companies {
		if singleCompany.Id.String() == companyId {
			singleCompany.Name = updatedCompany.Name
			companies = append(companies[:i], singleCompany)
			json.NewEncoder(w).Encode(singleCompany)
			return
		}
	}
}

func deleteCompany(w http.ResponseWriter, r *http.Request){
	companyId := mux.Vars(r)["id"]
	for i, singleCompany := range companies {
		if singleCompany.Id.String() == companyId {
			companies = append(companies[:i], companies[i +1:]...)
			fmt.Fprintf(w, "The company with ID %v has been deleted", companyId)
		}
	}
}

type Company struct {
	Id   guid.Guid `json:"Id"`
	Name string    `json:"Name"`
}

type allCompanies []Company

var companies = allCompanies{
	{
		Name: "Shirwood's Kitchen",
	},
}

package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type jsonStruct struct {
	Url string
	Name string
	Date string `json:",omitempty"`
	Description string `json:",omitempty"`
}

func openFile() []byte {
	inputJSON, err := ioutil.ReadFile("data/output.json")
	if err != nil {
		fmt.Println("Error opening file:\n")
		fmt.Println(err)
		return make([]byte, 0)
	}
	return inputJSON
}

func getData(w http.ResponseWriter, r *http.Request) {
	inputJSON := openFile()
	jsonMapping := make([]jsonStruct,0)
	errUnmarshal := json.Unmarshal([]byte(inputJSON), &jsonMapping)
	if errUnmarshal != nil {
		fmt.Println("Error Unmarshalling JSON:\n")
		fmt.Println(errUnmarshal)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(jsonMapping)

}

func getDataByName(w http.ResponseWriter, r *http.Request) {
	inputJSON := openFile()
	jsonMapping := make([]jsonStruct,0)
	errUnmarshal := json.Unmarshal([]byte(inputJSON), &jsonMapping)
	if errUnmarshal != nil {
		fmt.Println("Error Unmarshalling JSON:\n")
		fmt.Println(errUnmarshal)
	}
	params := mux.Vars(r) //Get All parameters
	expectedName := params["name"]
	w.Header().Set("Content-Type","application/json")
	for _, goItem := range jsonMapping {
		if goItem.Name == expectedName {
			json.NewEncoder(w).Encode(goItem)
			return
		}
	}

	return
}

func getDataByDate(w http.ResponseWriter, r *http.Request) {
	inputJSON := openFile()
	jsonMapping := make([]jsonStruct,0)
	errUnmarshal := json.Unmarshal([]byte(inputJSON), &jsonMapping)
	if errUnmarshal != nil {
		fmt.Println("Error Unmarshalling JSON:\n")
		fmt.Println(errUnmarshal)
	}
	params := mux.Vars(r) //Get All parameters
	expectedDate := params["date"]
	w.Header().Set("Content-Type","application/json")
	for _, goItem := range jsonMapping {
		if goItem.Date == expectedDate {
			json.NewEncoder(w).Encode(goItem)
			return
		}
	}
	return
}

func main() {
	muxRouter := mux.NewRouter()	
	muxRouter.HandleFunc("/goapi/getdata", getData).Methods("GET")
	muxRouter.HandleFunc("/goapi/getdata/{name}", getDataByName).Methods("GET")
	muxRouter.HandleFunc("/goapi/getdata/bydate/{date}", getDataByDate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}

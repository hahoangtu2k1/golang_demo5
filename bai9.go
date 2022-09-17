package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type useData struct {
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func bai9() {
	r := mux.NewRouter()
	r.HandleFunc("/get", GetData).Methods("GET")
	r.HandleFunc("/post", PostData).Methods("POST")
	log.Println("Server is running...")
	http.ListenAndServe(":1407", r)
}
func GetData(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Printf("Error getting data %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing data %s", err.Error())
	}
	var getData []useData
	err = json.Unmarshal(body, &getData)
	if err != nil {
		log.Printf("Error parsing json %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(getData)
	if err != nil {
		log.Printf("Error encoding json %s", err.Error())
	}
}
func PostData(w http.ResponseWriter, r *http.Request) {
	postData := useData{
		UserId: 123456,
		Id:     987654,
		Title:  "Post Data",
		Body:   "post data successfully",
	}

	data, _ := json.Marshal(postData)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing data %s", err.Error())
	}
	var posts useData
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Error parsing json %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Printf("Error encoding json %s", err.Error())
	}
}

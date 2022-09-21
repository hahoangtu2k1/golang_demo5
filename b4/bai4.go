package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hahoangtu2k1/b5"
	"github.com/rs/xid"
)

var engine = b5.Connect()

func bai4() {
	r := mux.NewRouter()

	r.HandleFunc("/GetUser", Get).Methods("GET")
	r.HandleFunc("/GetUser/{id}", GetById).Methods("GET")
	r.HandleFunc("/PostUser", Post).Methods("POST")
	r.HandleFunc("/DeleteUser/{id}", DeleteById).Methods("DELETE")
	log.Println("Server is running...")
	http.ListenAndServe(":3000", r)
}

func Get(w http.ResponseWriter, r *http.Request) {
	var engine = b5.Connect()
	var user []b5.UserPartner
	// engine.ShowSQL()
	err := engine.Table("user_partner").Find(&user)
	if err != nil {
		log.Print(err)
	}

	log.Println(user)

}
func GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user b5.UserPartner
	exist, _ := engine.Table("user_partner").Where("id = ?", id).Exist(&b5.UserPartner{
		Id: id,
	})
	if exist {
		_, err := engine.Table("user_partner").Where("id = ?", id).Get(&user)
		if err != nil {
			log.Println(err)
		}
		log.Println(user)
	} else {
		log.Println("Error")
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var user = b5.UserPartner{
		Id:          xid.New().String(),
		UserId:      xid.New().String(),
		PartnerId:   xid.New().String(),
		AliasUserId: xid.New().String(),
		Apps:        map[string]int64{"name": 7},
		Phone:       "0321456987",
		Created:     9876254,
		Updated_at:  4564987,
	}
	_, err := engine.Table("user_partner").Insert(&user)
	if err != nil {
		log.Print(err)
	}
	log.Print("insert successfully")

}
func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user b5.UserPartner
	exist, _ := engine.Table("user_partner").Where("id = ?", id).Exist(&b5.UserPartner{
		Id: id,
	})
	if exist {
		_, err := engine.Table("user_partner").Where("id = ?", id).Delete(&user)
		if err != nil {
			log.Println(err)
		}
		log.Println("DELETE done")
	} else {
		log.Println("Error")
	}
}

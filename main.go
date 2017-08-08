package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"encoding/json"
	"strconv"
)

type Foco struct {
	Id	int    `json:"id"`
	Estado	int    `json:"estado"`
	Descripcion string	`json:"descripcion"`
}

var Focos = make(map[string]Foco)
var id int

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/focos",GetFocosHandle).Methods("GET")
	r.HandleFunc("/focos",PostFocosHandle).Methods("POST")
	r.HandleFunc("/focos/{id}",PutFocosHandle).Methods("PUT")
	r.HandleFunc("/focos/{id}",DeleteFocosHandle).Methods("DELETE")
	server := &http.Server{
		Addr:	":5000",
		Handler:	r,
		ReadHeaderTimeout:	10*time.Second,
		WriteTimeout:	10*time.Second,
		MaxHeaderBytes:	1<<20,
	}
	println("Listening in localhost:5000")
	server.ListenAndServe()
}

func DeleteFocosHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k :=vars["id"]
	if _ , ok := Focos[k]; ok {
		delete(Focos,k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func PutFocosHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k :=vars["id"]
	var focoU Foco
	json.NewDecoder(r.Body).Decode(&focoU)
	if _ , ok := Focos[k]; ok {
		delete(Focos,k)
		Focos[k] = focoU
	}
	w.WriteHeader(http.StatusNoContent)
}

func PostFocosHandle(w http.ResponseWriter, r *http.Request) {
	var foco Foco
	json.NewDecoder(r.Body).Decode(&foco)
	id++
	k:=strconv.Itoa(id)
	Focos[k] = foco
	w.Header().Set("Content-Type","application/json")
	j,_ := json.Marshal(foco)
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetFocosHandle(w http.ResponseWriter, r *http.Request) {
	var focos []Foco
	for _,v := range Focos {
		focos = append(focos,v)
	}
	w.Header().Set("Content-Type","application/json")
	j,_ := json.Marshal(focos)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/michaelklishin/rabbit-hole"
	"net/http"
	"reflect"
	"strings"
)

func accessControl(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		f(w, r)
	}
}

func queue_detailsfetcher(w http.ResponseWriter, r *http.Request){
	rmqc, _ := rabbithole.NewClient("", "", "")
	file :=  strings.TrimPrefix(r.RequestURI, "/details/")
	q, err := rmqc.GetQueue("/", file)
	fmt.Println(q)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(q))
	b, err := json.Marshal(q)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "close")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func main() {

	r := mux.NewRouter()
	root_subrouter := r.PathPrefix("/").Subrouter()
	root_subrouter.Methods("GET").Name("Health").Path("/healthz").Handler(accessControl(func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	}))


	fetch_subrouter := r.PathPrefix("/details").Subrouter()
	fetch_subrouter.HandleFunc("/{queue_name}", queue_detailsfetcher)

	http.ListenAndServe(":8081", r)
}
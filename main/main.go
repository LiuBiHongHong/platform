package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/liubihonghong/platform"
	"github.com/rs/cors"
)

func GetAllAppHandler(w http.ResponseWriter, r *http.Request) {
	apps, err := platform.GetAllApp()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(apps)
	}
}

func GetAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	a, err := platform.GetApp(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(a)
	}
}

func StartAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	a, err := platform.StartApp(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(a)
	}
}

func DeleteAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	err := platform.DeleteApp(id)
	if err != nil {
		log.Println(err)
	}
}

func DownloadAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	a, err := platform.DownloadApp(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(a)
	}
}

func GetAllItemHandler(w http.ResponseWriter, r *http.Request) {
	items, err := platform.GetAllItem()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(items)
	}
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i, err := platform.GetItem(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(i)
	}
}

func GetAllServiceHandler(w http.ResponseWriter, r *http.Request) {}

func GetServiceHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app", GetAllAppHandler).Methods("GET")
	router.HandleFunc("/app/{id}", GetAppHandler).Methods("GET")
	router.HandleFunc("/app/start/{id}", StartAppHandler).Methods("GET")
	router.HandleFunc("/app/delete/{id}", DeleteAppHandler).Methods("DELETE")
	router.HandleFunc("/app/download/{id}", DownloadAppHandler).Methods("GET")
	router.HandleFunc("/item", GetAllItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", GetItemHandler).Methods("GET")
	router.HandleFunc("/service", GetAllServiceHandler).Methods("GET")
	router.HandleFunc("/service/{id}", GetServiceHandler).Methods("GET")
	router.HandleFunc("/service/delete/{id}", DeleteServiceHandler).Methods("DELETE")

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		platform.Test()
	}).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"X-Xsrf-Token", "Content-Type"},
	}).Handler(router)
	log.Fatal(http.ListenAndServe(":12000", handler))
}

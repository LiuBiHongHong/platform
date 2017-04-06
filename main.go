package main

import (
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"platform/platformlib"
)

func GetAllAppHandler(w http.ResponseWriter, r *http.Request) {}

func GetAppHandler(w http.ResponseWriter, r *http.Request) {}

func StartAppHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteAppHandler(w http.ResponseWriter, r *http.Request) {}

func DownloadAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	platformlib.DownloadApp(id)
}

func GetAllItemHandler(w http.ResponseWriter, r *http.Request) {}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {}

func GetAllServiceHandler(w http.ResponseWriter, r *http.Request) {}

func GetServiceHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app", GetAllAppHandler).Methods("GET")
	router.HandleFunc("/app/{id}", GetAppHandler).Methods("GET")
	router.HandleFunc("/startapp", StartAppHandler).Methods("POST")
	router.HandleFunc("/deleteapp", DeleteAppHandler).Methods("GET")
	router.HandleFunc("/downloadapp/{id}", DownloadAppHandler).Methods("GET")
	router.HandleFunc("/item", GetAllItemHandler).Methods("GET")
	router.HandleFunc("/item/{}", GetItemHandler).Methods("GET")
	router.HandleFunc("/service", GetAllServiceHandler).Methods("GET")
	router.HandleFunc("/service/{id}", GetServiceHandler).Methods("GET")
	router.HandleFunc("/deleteservice", DeleteServiceHandler).Methods("DELETE")

	// Cross-origin resource sharing
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"X-Xsrf-Token", "Content-Type"},
	}).Handler(router)
	log.Fatal(http.ListenAndServe(":12000", handler))
}

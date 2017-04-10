package main

import (
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"platform/platformlib"
)

func GetAllAppHandler(w http.ResponseWriter, r *http.Request) {
	_, err := platformlib.GetAllApp()
	if err != nil {
		log.Println(err)
	}
}

func GetAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	_, err := platformlib.GetApp(id)
	if err != nil {
		log.Println(err)
	}
}

func StartAppHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	err := platformlib.DeleteApp(id)
	if err != nil {
		log.Println(err)
	}
}

func DownloadAppHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	_, err := platformlib.DownloadApp(id)
	if err != nil {
		log.Println(err)
	}
}

func GetAllItemHandler(w http.ResponseWriter, r *http.Request) {
	_, err := platformlib.GetAllItem()
	if err != nil {
		log.Println(err)
	}
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	_, err := platformlib.GetItem(id)
	if err != nil {
		log.Println(err)
	}
}

func GetAllServiceHandler(w http.ResponseWriter, r *http.Request) {}

func GetServiceHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app", GetAllAppHandler).Methods("GET")
	router.HandleFunc("/app/{id}", GetAppHandler).Methods("GET")
	router.HandleFunc("/startapp", StartAppHandler).Methods("POST")
	router.HandleFunc("/deleteapp/{id}", DeleteAppHandler).Methods("DELETE")
	router.HandleFunc("/downloadapp/{id}", DownloadAppHandler).Methods("GET")
	router.HandleFunc("/item", GetAllItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", GetItemHandler).Methods("GET")
	router.HandleFunc("/service", GetAllServiceHandler).Methods("GET")
	router.HandleFunc("/service/{id}", GetServiceHandler).Methods("GET")
	router.HandleFunc("/deleteservice/{id}", DeleteServiceHandler).Methods("DELETE")

	// Cross-origin resource sharing
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"X-Xsrf-Token", "Content-Type"},
	}).Handler(router)
	log.Fatal(http.ListenAndServe(":12000", handler))
}

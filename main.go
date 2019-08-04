package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller"
	"./kategoriController"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/Item", controller.ShowAllItems).Methods("GET")
	router.HandleFunc("/api/Item", controller.InsertItemData).Methods("POST")
	router.HandleFunc("/api/Item/{id:[0-9]+}", controller.UpdateItemData).Methods("PUT")
	router.HandleFunc("/api/Item/{id:[0-9]+}", controller.DeleteItemData).Methods("DELETE")

	router.HandleFunc("/api/Kategori", kategoriController.ShowAllKategori).Methods("GET")
	router.HandleFunc("/api/Kategori", kategoriController.InsertKategoriData).Methods("POST")
	router.HandleFunc("/api/Kategori/{id:[0-9]+}", kategoriController.UpdateKategoriData).Methods("PUT")
	router.HandleFunc("/api/Kategori/{id:[0-9]+}", kategoriController.DeleteKategoriData).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")

	// define the port and load routes
	log.Fatal(http.ListenAndServe(":8080", router))

}

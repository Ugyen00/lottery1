package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/controller"
)

func InitRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/home", controller.HomeHandler)
	router.HandleFunc("/ticket", controller.BuyTicket).Methods("POST")
	router.HandleFunc("/ticket/{tid}", controller.DeleteTik).Methods("DELETE")
	router.HandleFunc("/ticket/{tid}", controller.UpdateTik).Methods("PUT")
	router.HandleFunc("/tickets", controller.GetAllTiks).Methods("GET")
	router.HandleFunc("/register",controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/login",controller.LoginHandler).Methods("POST")
	router.HandleFunc("/logout",controller.LogoutHandler).Methods("GET")

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application Running on prot 3003...")
	log.Fatal(http.ListenAndServe(":3003", router))
}

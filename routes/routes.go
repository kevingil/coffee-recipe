package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kevingil/coffee-recipe/coffeeapp"
)

func InitRouter() {
	r := mux.NewRouter()

	// Espresso App

	r.HandleFunc("/", coffeeapp.CoffeeApp).Methods("GET")
	r.HandleFunc("/api/stream-recipe", coffeeapp.StreamRecipe).Methods("POST")
	r.HandleFunc("/api/stream-recipe", coffeeapp.StreamRecipe).Methods("GET")

	//Files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Printf("Your app is running on port %s.", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

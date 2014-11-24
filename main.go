package main

import (
	"github.com/codegangsta/negroni"
	"github.com/go-gis/index-backend/handlers/ellipses"
	. "github.com/go-gis/index-backend/middlewares"
	"github.com/gorilla/mux"
	"os"
)

func main() {

	n := negroni.Classic()

	n.Use(MongoMiddleware())
	n.Use(negroni.HandlerFunc(CheckPostRequest))
	n.Use(NewAuthChecker())
	router := mux.NewRouter()

	router.HandleFunc("/ellipse/all", ellipses.AllEllipse).Methods("GET")
	router.HandleFunc("/ellipse/create", ellipses.CreateEllipse).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"github.com/codegangsta/negroni"
	"github.com/go-gis/index-backend/handlers/ellipses"
	"github.com/go-gis/index-backend/middlewares"
	"github.com/go-gis/index-backend/middlewares/auth"
	"github.com/go-gis/index-backend/middlewares/mongo"
	"github.com/gorilla/mux"
	"os"
)

func main() {

	n := negroni.Classic()

	n.Use(mongo.MongoMiddleware())
	n.Use(negroni.HandlerFunc(middlewares.CheckPostRequest))

	router := mux.NewRouter()

	router.HandleFunc("/ellipse/all", ellipses.AllEllipse).Methods("GET")
	router.HandleFunc("/ellipse/create", auth.SecureHandleFunc(ellipses.CreateEllipse)).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"github.com/codegangsta/negroni"
	"github.com/go-gis/index-backend/handlers/ellipsoid"
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

	router.HandleFunc("/ellipsoid/all", ellipsoid.AllEllipsoid).Methods("GET")
	router.HandleFunc("/ellipsoid/create", auth.SecureHandleFunc(ellipsoid.CreateEllipsoid)).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

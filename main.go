package main

import (
	"github.com/codegangsta/negroni"
	"github.com/go-gis/datex/handlers/ellipsoid"
	"github.com/go-gis/datex/handlers/meridian"
	"github.com/go-gis/datex/handlers/unit"
	"github.com/go-gis/datex/middlewares"
	"github.com/go-gis/datex/middlewares/auth"
	"github.com/go-gis/datex/middlewares/mongo"
	"github.com/gorilla/mux"
	"os"
)

func main() {

	n := negroni.Classic()

	n.Use(mongo.MongoMiddleware())
	n.Use(negroni.HandlerFunc(middlewares.CheckPostRequest))

	router := mux.NewRouter()

	router.HandleFunc("/ellipsoid/all", mongo.All(&ellipsoid.Ellipsoid{})).Methods("GET")
	router.HandleFunc("/ellipsoid/create", auth.SecureHandleFunc(mongo.Create(&ellipsoid.Ellipsoid{}))).Methods("POST")

	router.HandleFunc("/unit/all", mongo.All(&unit.Unit{})).Methods("GET")
	router.HandleFunc("/unit/create", auth.SecureHandleFunc(mongo.Create(&unit.Unit{}))).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

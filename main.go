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
	"github.com/meatballhat/negroni-logrus"
	"os"
)

func main() {

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.Use(mongo.MongoMiddleware())
	n.Use(negroni.HandlerFunc(middlewares.CheckPostRequest))

	router := mux.NewRouter()

	router.HandleFunc("/ellipsoid/all", mongo.All(&ellipsoid.Ellipsoid{})).Methods("GET")
	router.HandleFunc("/ellipsoid/create", auth.SecureHandleFunc(mongo.Create(&ellipsoid.Ellipsoid{}))).Methods("POST")

	router.HandleFunc("/unit/all", mongo.All(&unit.Unit{})).Methods("GET")
	router.HandleFunc("/unit/create", auth.SecureHandleFunc(mongo.Create(&unit.Unit{}))).Methods("POST")

	router.HandleFunc("/meridian/all", mongo.All(&meridian.Meridian{})).Methods("GET")
	router.HandleFunc("/meridian/create", auth.SecureHandleFunc(mongo.Create(&meridian.Meridian{}))).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

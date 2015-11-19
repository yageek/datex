package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/yageek/datex/handlers/ellipsoid"
	"github.com/yageek/datex/handlers/meridian"
	"github.com/yageek/datex/handlers/unit"
	"github.com/yageek/datex/middlewares"
	"github.com/yageek/datex/middlewares/auth"
	"github.com/yageek/datex/middlewares/mongo"
	"os"
)

func main() {

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.Use(mongo.MongoMiddleware())
	n.Use(negroni.HandlerFunc(middlewares.CheckPostRequest))

	router := mux.NewRouter()

	router.HandleFunc("/ellipsoid/all", mongo.All(&ellipsoid.Ellipsoid{})).Methods("GET")
	router.HandleFunc("/ellipsoid/create", auth.SecureHandleFunc(mongo.Create((*ellipsoid.Ellipsoid)(nil)))).Methods("POST")

	router.HandleFunc("/unit/all", mongo.All(&unit.Unit{})).Methods("GET")
	router.HandleFunc("/unit/create", auth.SecureHandleFunc(mongo.Create((*unit.Unit)(nil)))).Methods("POST")

	router.HandleFunc("/meridian/all", mongo.All(&meridian.Meridian{})).Methods("GET")
	router.HandleFunc("/meridian/create", auth.SecureHandleFunc(mongo.Create((*meridian.Meridian)(nil)))).Methods("POST")

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

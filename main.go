package main

import (
	"github.com/codegangsta/negroni"
	"github.com/go-gis/index-backend/database"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"os"
)

var urender *render.Render

func init() {
	urender = render.New(render.Options{})
}

func main() {

	n := negroni.Classic()
	n.Use(database.MongoMiddleware())

	router := mux.NewRouter()

	n.UseHandler(router)

	n.Run(":" + os.Getenv("PORT"))
}

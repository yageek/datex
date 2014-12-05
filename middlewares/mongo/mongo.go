//Middleware handling mgo sessions
package mongo

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/go-gis/datex/middlewares"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"reflect"
)

var mongoSession *mgo.Session

type key int

const db key = 0

const DatabaseName = "opengis_index"

type IndexObject interface {
	CollectionName() string
	ValidCreation() (bool, string)
}

// Add the mgo database to the request context
func SetDb(r *http.Request, val *mgo.Database) {
	context.Set(r, db, val)
}

// Get the mgo database from the request context
func GetDb(r *http.Request) *mgo.Database {
	if rv := context.Get(r, db); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}

// Returns negroni middleware mapping a mgo database
// from a cloned mgo session.
func MongoMiddleware() negroni.HandlerFunc {

	url := os.Getenv("MONGODB_URL")
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqSession := session.Clone()
		defer reqSession.Close()
		db := reqSession.DB(DatabaseName)
		SetDb(r, db)
		next(rw, r)
	})
}

// Get the mgo collection from request context by its name
func Collection(r *http.Request, name string) *mgo.Collection {
	return GetDb(r).C(name)
}

func All(indexObject IndexObject) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c := Collection(r, indexObject.CollectionName())

		var results []interface{}
		err := c.Find(bson.M{}).All(&results)

		if err != nil {
			log.Println("MONGO error:", err)
			http.Error(w, "Could not retrieve object", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if len(results) == 0 {
			w.Write([]byte("[]"))
		} else {
			data, _ := json.Marshal(results)
			w.Write(data)
		}

	}
}

func Create(i IndexObject) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		t := reflect.TypeOf(i).Elem()
		indexObject := reflect.New(t).Interface().(IndexObject)

		data := middlewares.GetData(r)

		err := json.Unmarshal(data, indexObject)

		if err != nil {
			log.Printf("Error by decoding JSON:%s\n", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		valid, message := indexObject.ValidCreation()

		if !valid {
			log.Println(message)
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		c := Collection(r, indexObject.CollectionName())

		err = c.Insert(indexObject)

		if err != nil {
			log.Printf("Error by inserting data: %s\n", err)
			http.Error(w, "DB error :(", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

//Middleware handling mgo sessions
package mongo

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
)

var mongoSession *mgo.Session

type key int

const db key = 0

const DatabaseName = "opengis_index"

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

//Manages handlers targeting ellipsoids
package ellipsoid

import (
	"encoding/json"
	. "github.com/go-gis/index-backend/middlewares"
	"github.com/go-gis/index-backend/middlewares/mongo"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

const EllipsoidCollectionName = "ellipsoids"

type Ellipsoid struct {
	Name          string  `json:"name" bson:"name"`
	SemiMajorAxis float64 `json:"a"  bson:"semi-major-axis"`
	Deprecated    bool    `json:"deprecated" bson:"deprecated"`

	EPSG              int     `json:"epsg,omitempty" bson:"epsg,omitempty"`
	Description       string  `json:"description,omitempty" bson:"description,omitempty"`
	Source            string  `json:"source,omitempty" bson:"source,omitempty"`
	SemiMinorAxis     float64 `json:"b,omitempty" bson:"semi-minor-axis,omitempty"`
	InverseFlattening float64 `json:"f,omitempty"  bson:"inverse-flattening,omitempty"`
}

// Returns all the ellipsoids presents in the database
func AllEllipsoid(w http.ResponseWriter, r *http.Request) {
	c := mongo.Collection(r, EllipsoidCollectionName)

	var results []Ellipsoid
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		http.Error(w, "Could not retrieve ellipse", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(results)

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

// Create an ellipsoid with the JSON provided in the request.
// The JSON should provide least have a name, a semi-major-axis, a semi-minor-axis
// or the inverse flattening. If the deprecated value is not provided, it is considered
// as false.
func CreateEllipsoid(w http.ResponseWriter, r *http.Request) {

	e := Ellipsoid{}
	data := GetData(r)

	err := json.Unmarshal(data, &e)

	if err != nil {
		log.Printf("Error by decoding JSON:%s\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if e.Name == "" || e.SemiMinorAxis == 0 || (e.SemiMinorAxis == 0 && e.InverseFlattening == 0) {
		message := "Ellipsoid name, semi-major-axis, inverse-flattening or semi-minor-axis are mandatory"
		log.Println(message)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	c := mongo.Collection(r, EllipsoidCollectionName)

	err = c.Insert(e)

	if err != nil {
		log.Printf("Error by inserting data: %s\n", err)
		http.Error(w, "DB error :(", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

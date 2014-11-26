package ellipses

import (
	"encoding/json"
	. "github.com/go-gis/index-backend/middlewares"
	"github.com/go-gis/index-backend/middlewares/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

const EllipseCollectionName = "ellipses"

type Ellipse struct {
	Description   string  `json:"description" bson:"description"`
	SemiMajorAxis float64 `json:"a"  bson:"major-axis"`
	Flattening    float64 `json:"f"  bson:"flattening"`
	Reference     string  `json:"ref" bson:"reference"`
}

func collection(r *http.Request) *mgo.Collection {
	return mongo.GetDb(r).C(EllipseCollectionName)
}

func AllEllipse(w http.ResponseWriter, r *http.Request) {
	c := collection(r)

	var results []Ellipse
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		http.Error(w, "Could not retrieve ellipse", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(results)

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func CreateEllipse(w http.ResponseWriter, r *http.Request) {

	e := Ellipse{}
	data := GetData(r)

	err := json.Unmarshal(data, &e)

	if err != nil {
		log.Printf("Error by decoding JSON:%s\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	c := collection(r)

	err = c.Insert(e)

	if err != nil {
		log.Printf("Error by inserting data: %s\n", err)
		http.Error(w, "DB error :(", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

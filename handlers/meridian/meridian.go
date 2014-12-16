package meridian

import (
	"math"
	"time"
)

const MeridianCollectionName = "prime_meridians"

type Meridian struct {
	Name              string    `json:"name" bson:"name"`
	EPSG              int       `json:"epsg" bson:"epsg"`
	RevisionDate      time.Time `json:"revision_date,omitempty" bson:"revision_date,omitempty"`
	Deprecated        bool      `json:"deprecated" bson:"deprecated"`
	Source            string    `json:"source,omitempty" bson:"source,omitempty"`
	Description       string    `json:"description,omitempty" bson:"description,omitempty"`
	GrenwichLongitude float64   `json:"greenwich_longitude" bson:"greenwich_longitude"`
}

func (m *Meridian) CollectionName() string {
	return MeridianCollectionName
}

func (m *Meridian) ValidCreation() (bool, string) {
	if m.Name == "" || m.GrenwichLongitude == math.NaN() {
		return false, "Name and Greenwich Longitude are mandatory"
	}

	return true, ""
}

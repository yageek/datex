package meridian

import (
	"math"
)

const MeridianCollectionName = "prime_merdian"

type Meridian struct {
	Name              string  `json:"name" bson:"name"`
	GrenwichLongitude float64 `json:"greenwich_longitude" bson:"greenwich_longitude"`
	EPSG              int     `json:"epsg,omitempty" bson:"epsg,omitempty"`
	Description       string  `json:"description,omitempty" bson:"description,omitempty"`
	Source            string  `json:"source,omitempty" bson:"source,omitempty"`
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

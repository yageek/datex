package datum

import (
	"github.com/go-gis/datex/handlers/ellipsoid"
	"github.com/go-gis/datex/handlers/meridian"
)

type DatumType int

const (
	DatumCollectionName           = "datum"
	Vertical            DatumType = iota
	Geodetic
	Engineering
)

type Datum struct {
	Name          string    `json:"name" bson:"name"`
	Deprecated    bool      `json:"deprecated" bson:"deprecated"`
	Ellipsoid     ellipsoid `json:"ellipsoid" bson:"ellipsoid"`
	Type          DatumType `json:"type" bson:"type"`
	PrimeMeridian meridian  `json:"prime_meridian" bson:"prime_meridian"`
	EPSG          int       `json:"epsg,omitempty" bson:"epsg,omitempty"`
	Epoch         int       `json:"epoch,omitempty" bson:"epoch,omitempty"`
	Description   string    `json:"description,omitempty" bson:"description,omitempty"`
	Source        string    `json:"source,omitempty" bson:"source,omitempty"`
}

func (d *Datum) CollectionName() string {
	return DatumCollectionName
}

func (m *Datum) ValidCreation() (bool, string) {
	return true, ""
}

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
	Name          string               `json:"name" bson:"name"`
	Deprecated    bool                 `json:"deprecated" bson:"deprecated"`
	Ellipsoid     *ellipsoid.Ellipsoid `json:"ellipsoid" bson:"ellipsoid"`
	Type          DatumType            `json:"type" bson:"type"`
	PrimeMeridian *meridian.Meridian   `json:"prime_meridian" bson:"prime_meridian"`
	EPSG          int                  `json:"epsg,omitempty" bson:"epsg,omitempty"`
	Epoch         int                  `json:"epoch,omitempty" bson:"epoch,omitempty"`
	Description   string               `json:"description,omitempty" bson:"description,omitempty"`
	Source        string               `json:"source,omitempty" bson:"source,omitempty"`
	RevisionDate  time.Time            `json:"revision_date,omitempty" bson:"revision_date,omitempty"`
}

func (d *Datum) CollectionName() string {
	return DatumCollectionName
}

func (d *Datum) ValidCreation() (bool, string) {
	if d.Name == "" || d.Ellipsoid == nil || d.PrimeMeridian == nil {
		return false, "Name, Ellipsoid and Prime meridian are mandatory"
	}
	return true, ""
}

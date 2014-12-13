//Manages handlers targeting ellipsoids
package ellipsoid

import (
	"time"
)

const EllipsoidCollectionName = "ellipsoids"

type Ellipsoid struct {
	Name              string    `json:"name" bson:"name"`
	EPSG              int       `json:"epsg" bson:"epsg"`
	RevisionDate      time.Time `json:"revision_date,omitempty" bson:"revision_date,omitempty"`
	Deprecated        bool      `json:"deprecated" bson:"deprecated"`
	Source            string    `json:"source,omitempty" bson:"source,omitempty"`
	Description       string    `json:"description,omitempty" bson:"description,omitempty"`
	SemiMajorAxis     float64   `json:"semi-major-axis"  bson:"semi-major-axis"`
	SemiMinorAxis     float64   `json:"semi-minor-axis,omitempty" bson:"semi-minor-axis,omitempty"`
	InverseFlattening float64   `json:"inverse-flattening,omitempty"  bson:"inverse-flattening,omitempty"`
}

func (e *Ellipsoid) CollectionName() string {
	return EllipsoidCollectionName
}

func (e *Ellipsoid) ValidCreation() (bool, string) {
	if e.Name == "" || e.SemiMajorAxis <= 0 || (e.SemiMinorAxis <= 0 && e.InverseFlattening <= 0) {
		return false, "Ellipsoid name, semi-major-axis, inverse-flattening or semi-minor-axis are mandatory"
	}

	return true, ""
}

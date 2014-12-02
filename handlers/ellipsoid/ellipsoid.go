//Manages handlers targeting ellipsoids
package ellipsoid

const EllipsoidCollectionName = "ellipsoids"

type Ellipsoid struct {
	Name          string  `json:"name" bson:"name"`
	SemiMajorAxis float64 `json:"a"  bson:"semi-major-axis"`
	Deprecated    bool    `json:"deprecated" bson:"deprecated"`

	EPSG              int     `json:"epsg,omitempty" bson:"epsg,omitempty"`
	Description       string  `json:"description,omitempty" bson:"description,omitempty"`
	Source            string  `json:"source,omitempty" bson:"source,omitempty"`
	SemiMinorAxis     float64 `json:"b,omitempty" bson:"semi-minor-axis,omitempty"`
	InverseFlattening float64 `json:"inv_f,omitempty"  bson:"inverse-flattening,omitempty"`
}

func (e *Ellipsoid) CollectionName() string {
	return EllipsoidCollectionName
}

func (e *Ellipsoid) ValidCreation() (bool, string) {
	if e.Name == "" || e.SemiMinorAxis <= 0 || (e.SemiMinorAxis <= 0 && e.InverseFlattening <= 0) {
		return false, "Ellipsoid name, semi-major-axis, inverse-flattening or semi-minor-axis are mandatory"
	}

	return true, ""
}

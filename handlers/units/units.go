package units

type UnitType int

const (
	UnitCollectionName = "ellipsoids"

	Length UnitType = iota
	Time
	Scale
	Angle
)

type Units struct {
	EPSG        int    `json:"epsg" bson:"epsg"`
	Name        string `json:"name" bson:"name"`
	Source      string `json:"source,omitempty" bson:"source,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

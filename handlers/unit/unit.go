package unit

type UnitType int

const (
	UnitCollectionName = "ellipsoids"

	Length UnitType = iota
	Time
	Scale
	Angle
)

type Unit struct {
	EPSG        int    `json:"epsg" bson:"epsg"`
	Name        string `json:"name" bson:"name"`
	Source      string `json:"source,omitempty" bson:"source,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

func (u *Unit) CollectionName() string {
	return UnitCollectionName
}

func (u *Unit) ValidCreation() (bool, string) {
	if u.EPSG == 0 || u.Name == "" {
		return false, "EPSG and Name is mandatory"
	} else {
		return true, ""
	}
}
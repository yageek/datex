package unit

import (
	"time"
)

type UnitType int

const (
	UnitCollectionName = "units"

	Length UnitType = iota
	Time
	Scale
	Angle
)

func (t UnitType) String() string {
	switch t {
	case Time:
		return "time"
	case Length:
		return "length"
	case Scale:
		return "scale"
	case Angle:
		return "angle"
	default:
		return "unknown"
	}
}

type Unit struct {
	EPSG         int       `json:"epsg" bson:"epsg"`
	Name         string    `json:"name" bson:"name"`
	Deprecated   bool      `json:"deprecated" bson:"deprecated"`
	Source       string    `json:"source,omitempty" bson:"source,omitempty"`
	Description  string    `json:"description,omitempty" bson:"description,omitempty"`
	Type         UnitType  `json:"type" bson:"type"`
	RevisionDate time.Time `json:"revision_date,omitempty" bson:"type,omitempty"`
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

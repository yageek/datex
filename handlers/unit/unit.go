package unit

import (
	"time"
)

type UnitType int

const (
	UnitCollectionName = "units"
)

type Unit struct {
	Name         string    `json:"name" bson:"name"`
	EPSG         int       `json:"epsg" bson:"epsg"`
	RevisionDate time.Time `json:"revision_date,omitempty" bson:"revision_date,omitempty"`
	Deprecated   bool      `json:"deprecated" bson:"deprecated"`
	Source       string    `json:"source,omitempty" bson:"source,omitempty"`
	Description  string    `json:"description,omitempty" bson:"description,omitempty"`
	Type         string    `json:"unit_type" bson:"unit_type"`
	FactorB      float64   `json:"factor_b" bson:"factor_b"`
	FactorC      float64   `json:"factor_c" bson:"factor_c"`
}

func (u *Unit) CollectionName() string {
	return UnitCollectionName
}

func (u *Unit) ValidCreation() (bool, string) {
	if u.EPSG == 0 || u.Name == "" || u.FactorB == 0 || u.FactorC == 0 {
		return false, "EPSG, Name, B and C Factor are mandatory"
	} else {
		return true, ""
	}
}

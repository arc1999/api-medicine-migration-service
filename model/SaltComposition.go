package model

type SaltComposition struct {
	SaltName string `json:"saltName"`
	Strength string `json:"strength"`
	Uom      string `json:"uom"`
	Atc      string `json:"atc"`
}

package model

import (
	"time"
)

type Medicine struct {
	ID                   int64 `json:"id,omitempty" bson:"_id,omitempty"`
	MedicineName         string             `json:"medicineName" bson:"medicineName"`
	ManufaturerName      string             ` json:"manufaturerName" bson:"manufaturerName"`
	MCode                string             `json:"mCode" bson:"mCode"` // This is Temporary As we don't have codes for all medicines;
	NdcCode              string             `json:"ndcCode" bson:"ndcCode"`
	Formulation          string             `json:"formulation" bson:"formulation"` // Different formulations in which drug exist
	DateCreated          time.Time          `json:"dateCreated" bson:"dateCreated"`
	DateUpdated          time.Time          ` json:"dateUpdated" bson:"dateUpdated"`
	CreatedBy            int64              ` json:"createdBy" bson:"createdBy"`
	UpdatedBy            int64              `json:"updatedBy" bson:"updatedBy"`
	CreatedByName        string             `json:"createdByName" bson:"createdByName"`
	UpdatedByName        string             `json:"updatedByName" bson:"updatedByName"`
	PrescriptionRequired bool               `json:"prescriptionRequired" bson:"prescriptionRequired"`
	SaltComposition      []SaltComposition  `json:"saltComposition" bson:"saltComposition"`
}

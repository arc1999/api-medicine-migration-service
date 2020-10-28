package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MedicineMongo struct {
	ID                  int64 `json:"id,omitempty" bson:"_id,omitempty"`
	MedicineName         string             `json:"medicineName" bson:"medicineName"`
	ManufaturerName      string             ` json:"manufaturerName" bson:"manufaturerName"`
	MCode                string             `json:"mCode" bson:"mCode"` // This is Temporary As we don't have codes for all medicines;
	Slug                 primitive.Binary   `json:"slug"`
	DateCreated          time.Time          `json:"dateCreated" bson:"dateCreated"`
	DateUpdated          time.Time          ` json:"dateUpdated" bson:"dateUpdated"`
	CreatedBy            int64              ` json:"createdBy" bson:"createdBy"`
	UpdatedBy            int64              `json:"updatedBy" bson:"updatedBy"`
	PrescriptionRequired bool               `json:"prescriptionRequired" bson:"prescriptionRequired"`
	SaltComposition      []SaltComposition  `json:"saltComposition" bson:"saltComposition"`
}

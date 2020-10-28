package transformer

import (
	"api-medicine-migration-service/model"
)

func Transform(mongomedicine []model.MedicineMongo) []model.Medicine {
	var medicines []model.Medicine
	for _, mmedicine := range mongomedicine {
		var medicine model.Medicine
		medicine.DateUpdated = mmedicine.DateUpdated
		medicine.ID = mmedicine.ID
		medicine.DateCreated = mmedicine.DateCreated
		medicine.UpdatedBy = mmedicine.UpdatedBy
		medicine.CreatedBy = mmedicine.CreatedBy
		medicine.MCode = mmedicine.MCode
		medicine.PrescriptionRequired = mmedicine.PrescriptionRequired
		medicine.SaltComposition = mmedicine.SaltComposition
		medicine.ManufaturerName=mmedicine.ManufaturerName
		medicine.MedicineName=mmedicine.MedicineName
		medicines = append(medicines, medicine)
	}
	return medicines
}

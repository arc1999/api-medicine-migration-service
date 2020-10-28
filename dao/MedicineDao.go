package dao

import (
	"api-medicine-migration-service/db"
	"api-medicine-migration-service/model"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MedicineDao struct {
}

func (d MedicineDao) Paginate(pagenumber int64, nperpage int64) ([]model.MedicineMongo, error) {
	log.Println(pagenumber, nperpage)
	findOptions := options.Find()
	findOptions.SetLimit(nperpage)
	findOptions.SetSort(bson.M{})
	findOptions.SetSkip(pagenumber)

	db := db.GetMongoDB()
	cur, err := db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var jobs []model.MedicineMongo
	for cur.Next(context.TODO()) {
		var job model.MedicineMongo
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	log.Println(len(jobs))
	return jobs, nil
}
func (d MedicineDao) GetCount() (int64, error) {
	db := db.GetMongoDB()
	return db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).CountDocuments(context.TODO(), bson.D{})
}
func (d MedicineDao) BulkInsert(Entity []model.Medicine) error {
	mongodb := db.GetMongoDB()
	var bulk []interface{}
	for _, item := range Entity {
		bulk = append(bulk, item)
	}
	_, err := mongodb.Collection(os.Getenv("NEW_MONGODB_COLLECTION")).InsertMany(context.TODO(), bulk)
	if err != nil {
		log.Printf("error in saving medicines")
		return err
	}
	return nil

}

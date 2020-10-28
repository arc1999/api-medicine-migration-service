package service

import (
	"api-medicine-migration-service/dao"
	"api-medicine-migration-service/transformer"
	"log"
	"os"
	"strconv"
)

var d dao.MedicineDao

type MedicineService struct {
}

func (is MedicineService) Migrate() {
	totaldoc, err := d.GetCount()
	if err != nil {
		log.Fatal(err)
	}
	perpage := os.Getenv("N_PER_PAGE")
	nperpage, err := strconv.ParseInt(perpage, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var i int64
	i = 0
	var noofpages = totaldoc / nperpage
	log.Println(totaldoc)
	log.Println(noofpages)
	log.Println(nperpage)
	for i < noofpages {
		mongomedicine, err := d.Paginate(i*nperpage, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		medicines := transformer.Transform(mongomedicine)

		err = d.BulkInsert(medicines)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	mongomedicine, err := d.Paginate(i*nperpage, totaldoc-(nperpage*(i)))
	if err != nil {
		log.Fatal(err)
	}
	medicines := transformer.Transform(mongomedicine)
	err = d.BulkInsert(medicines)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import "api-medicine-migration-service/service"

var s service.MedicineService

func main() {
	s.Migrate()
}

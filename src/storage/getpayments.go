package storage

import (
	"models"
	"encoding/json"
)

func GetPayments() string{
	var db = GetDb()
	var vehicleSlice []models.Vehicle
	results, err := db.Query("SELECT model,plate,capacity,vin FROM fuel_payment;")
	if err != nil{
		return "Database error"
	}
	defer results.Close()
	for results.Next() {
		var model string
		var plate string
		var capacity float32
		var vin string
		var v models.Vehicle
		err = results.Scan(&model, &plate, &capacity, &vin)
		if err != nil{
			continue
		}
		v.SetAttributes(model, plate, capacity, vin)
		vehicleSlice = append(vehicleSlice, v)
	}
	data ,_ := json.MarshalIndent(vehicleSlice,"","\t")
	return string(data)
}

//select amount,name,surname,plate from fuel_payment inner join fuel_driver on fuel_payment.driver_id = fuel_driver.id and fuel_driver.id=1 inner join fuel_vehicle on fuel_payment.vehicle_id = fuel_vehicle.id;


package storage

import (
	"models"
	"encoding/json"
)

func GetVehicles() string{
	var db = GetDb()
	var vehicleSlice []models.Vehicle
	results, err := db.Query("SELECT model,plate,capacity,vin FROM fuel_vehicle;")
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

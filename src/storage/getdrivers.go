package storage

import (
	"models"
	"encoding/json"
)

func GetDrivers() string{
	var db = GetDb()
	var driverSlice []models.Driver
	results, err := db.Query("SELECT name,surname,photo FROM fuel_driver;")
	if err != nil{
		return "Database error"
	}
	defer results.Close()
	for results.Next() {
		var name string
		var surname string
		var photo string
		var d models.Driver
		err = results.Scan(&name, &surname, &photo)
		if err != nil{
			continue
		}
		d.SetAttributes(name, surname, photo)
		driverSlice = append(driverSlice, d)
	}
	data ,_ := json.MarshalIndent(driverSlice,"","\t")
	return string(data)
}

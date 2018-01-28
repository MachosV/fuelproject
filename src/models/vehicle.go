package models

type Vehicle struct{
	Model string `json:"model"`
	Plate string `json:"plate"`
	Capacity float32 `json:"capacity"`
	Vin string `json:"vin"`
}

func (v *Vehicle) SetAttributes(model string, plate string, capacity float32, vin string){
	v.Model = model
	v.Plate = plate
	v.Capacity = capacity
	v.Vin = vin
}
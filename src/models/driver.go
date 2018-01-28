package models

type Driver struct{
	Name string `json:"name"`
	Surname string `json:"surname"`
	Photo string `json:"photo"`
}

func (d *Driver) SetAttributes(name string, surname string, photo string){
	d.Name = name
	d.Surname = surname
	d.Photo = photo
}
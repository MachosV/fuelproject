package models

type Driver struct{
	Name string `json:"name"`
	Surname string `json:"surname"`
	Photo string `json:"photo"`
}

func (d *Driver) GetName()string{
	return d.Name
}

func (d *Driver) GetSurname()string{
	return d.Surname
}
func (d *Driver) GetPhoto()string{
	return d.Photo
}

func (d *Driver) SetName(name string){
	d.Name = name
}

func (d *Driver) SetSurname(surname string) {
	d.Surname = surname
}

func (d *Driver) SetPhoto(photo string) {
	d.Photo = photo
}

func (d *Driver) SetAttributes(name string, surname string, photo string){
	d.Name = name
	d.Surname = surname
	d.Photo = photo
}
package models

type Manufacturer struct {
	id                 int    `json:"id"`
	Pharmacy_license   string `json:"pharmacy_license"`
	Pharmacist_license string `json:"pharmacist_license"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	Po_box             string `json:"po_box"`
	Email              string `json:"email"`
	Phone_number       string `json:"phone_number"`
}

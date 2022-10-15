package models

import "time"

type Drug struct {
	id                int       `json:"id"`
	Drug_name         string    `json:"drug_name"`
	Drug_generic_name string    `json:"drug_generic_name"`
	Dosage            int       `json:"dosage"`
	Date_enrolled     time.Time `json:"date_enrolled"`
	Expiry_date       time.Time `json:"expiry_date"`
	Nafdac            string    `json:"nafdac"`
	Qr_code           string    `json:"qr_code"`
	Manufacturers_id  int       `json:"manufacturers_id"`
}

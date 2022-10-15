package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/bits-his/drugCipherMiddleware/models"
)

type MySqlRepository struct {
	db *sql.DB
}

func NewSingleStoreRepository(db *sql.DB) *MySqlRepository {

	return &MySqlRepository{
		db: db,
	}
}

func (repo MySqlRepository) Drugs() ([]models.Drug, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), aLL_DRUGS)
	if err != nil {
		return []models.Drug{}, err
	}
	defer stmt.Close()
	var drugs []models.Drug

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return []models.Drug{}, err
	}
	{

		for rows.Next() {
			var drug models.Drug
			err := rows.Scan(&drug.Id, &drug.Drug_name, &drug.Drug_generic_name, &drug.Dosage,
				&drug.Date_enrolled, &drug.Expiry_date, &drug.Nafdac,
				&drug.Qr_code, &drug.Manufacturers_id)

			if err != nil {
				log.Fatal(err)
			}
			drugs = append(drugs, drug)

		}
	}
	return drugs, nil
}

func (repo MySqlRepository) Manufacturers() ([]models.Manufacturer, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), aLL_DRUGS)
	if err != nil {
		return []models.Manufacturer{}, err
	}
	defer stmt.Close()
	var manufacturers []models.Manufacturer

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return []models.Manufacturer{}, err
	}
	{

		for rows.Next() {
			var manufacturer models.Manufacturer
			err := rows.Scan(&manufacturer.Id, &manufacturer.Pharmacy_license, &manufacturer.Pharmacist_license,
				&manufacturer.Name, &manufacturer.Address, &manufacturer.Po_box,
				&manufacturer.Email, &manufacturer.Phone_number)

			if err != nil {
				log.Fatal(err)
			}
			manufacturers = append(manufacturers, manufacturer)

		}
	}
	return manufacturers, nil
}
func (repo MySqlRepository) GET_MANUFACTURER(id int) (models.Drug, error) {
	return models.Drug{}, nil
}
func (repo MySqlRepository) GET_DRUG(id int) (models.Manufacturer, error) {
	return models.Manufacturer{}, nil
}
func (repo MySqlRepository) UPDATE_MANUFACTURER(m *models.Manufacturer) (*models.Manufacturer, error) {
	return m, nil
}
func (repo MySqlRepository) UPDATE_DRUG(d *models.Drug) (*models.Drug, error) {
	return d, nil
}

func (repo MySqlRepository) MIGRATE_ALL() error {
	_, err := repo.db.ExecContext(context.Background(), cREATE_TABLE_MANUFACTURERS)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = repo.db.ExecContext(context.Background(), cREATE_TABLE_DRUGS)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

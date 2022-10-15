package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/bits-his/drugCipherMiddleware/models"
)

var ctx context.Context

type MySqlRepository struct {
	db *sql.DB
}

func NewSingleStoreRepository(db *sql.DB) *MySqlRepository {

	return &MySqlRepository{
		db: db,
	}
}

func (repo MySqlRepository) Drugs() ([]models.Drug, error) {

	return []models.Drug{}, nil
}

func (repo MySqlRepository) Manufacturers() ([]models.Manufacturer, error) {
	return []models.Manufacturer{}, nil
}
func (repo MySqlRepository) GET_MANUFACTURER(id int) (models.Drug, error) {
	return models.Drug{}, nil
}
func (repo MySqlRepository) GET_DRUG(id int) (models.Manufacturer, error) {
	return models.Manufacturer{}, nil
}
func (repo MySqlRepository) UPDATE_MANUFACTURER(m *models.Manufacturer) {}
func (repo MySqlRepository) UPDATE_DRUG(d *models.Drug)                 {}

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

package db

import "github.com/bits-his/drugCipherMiddleware/models"

type Repository interface {
	Drugs() ([]models.Drug, error)
	Manufacturers() ([]models.Manufacturer, error)
	GET_MANUFACTURER(id int) (models.Drug, error)
	GET_DRUG(id int) (models.Manufacturer, error)
	UPDATE_MANUFACTURER(m *models.Manufacturer) (*models.Manufacturer, error)
	UPDATE_DRUG(d *models.Drug) (*models.Drug, error)
	MIGRATE_ALL() error
}

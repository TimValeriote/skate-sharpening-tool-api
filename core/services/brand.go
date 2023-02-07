package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type brandStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func BrandStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *brandStore {
	return &brandStore{
		database: db,
		log:      log,
	}
}

func (store *brandStore) GetAllBrands() ([]models.BrandStruct, error) {
	sql := `SELECT id, name, short_name, is_skate, is_steel, is_holder FROM brands`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "brandStore::GetAllBrands - Failed to prepare GetAllBrands SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getBrandsFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *brandStore) GetBrandById(brandId int) ([]models.BrandStruct, error) {
	sql := `SELECT id, name, short_name, is_skate, is_steel, is_holder FROM brands WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "brandStore::GetBrandById - Failed to prepare GetBrandById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(brandId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := make([]models.BrandStruct, 0)
	for rows.Next() {
		var brand models.BrandStruct
		err = rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.ShortName,
			&brand.IsSkate,
			&brand.IsSteel,
			&brand.IsHolder,
		)
		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return brands, nil
}

func getBrandsFromQuery(query *sql.Stmt) ([]models.BrandStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := make([]models.BrandStruct, 0)
	for rows.Next() {
		var brand models.BrandStruct
		err = rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.ShortName,
			&brand.IsSkate,
			&brand.IsSteel,
			&brand.IsHolder,
		)
		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return brands, nil
}

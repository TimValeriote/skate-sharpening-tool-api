package core

import (
	"database/sql"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type skateStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func SkateStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *skateStore {
	return &skateStore{
		database: db,
		log:      log,
	}
}

func (store *skateStore) GetAllSkates() ([]models.SkateStruct, error) {
	sql := `SELECT 
					skates.id,

					skate_model.id,
					skate_model.name, 
					skate_model.alias,

					skate_brand.id,
					skate_brand.name, 
					skate_brand.short_name, 
					skate_brand.is_skate, 
					skate_brand.is_steel, 
					skate_brand.is_holder
			FROM skates
			INNER JOIN model AS skate_model ON skates.model_id = skate_model.id
			INNER JOIN brands AS skate_brand ON skates.brand_id = skate_brand.id`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "skateStore::GetAllSkates - Failed to prepare GetAllSkates SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getSkatesFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *skateStore) GetSkateById(skateId int) ([]models.SkateStruct, error) {
	sql := `SELECT 
					skates.id,

					skate_model.id,
					skate_model.name, 
					skate_model.alias,

					skate_brand.id,
					skate_brand.name, 
					skate_brand.short_name, 
					skate_brand.is_skate, 
					skate_brand.is_steel, 
					skate_brand.is_holder
			FROM skates
			INNER JOIN model AS skate_model ON skates.model_id = skate_model.id
			INNER JOIN brands AS skate_brand ON skates.brand_id = skate_brand.id
			WHERE skates.id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "skateStore::GetSkateBySkateId - Failed to prepare GetSkateBySkateId SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(skateId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	skatesArray := make([]models.SkateStruct, 0)
	for rows.Next() {
		var skate models.SkateStruct

		var model models.ModelStruct

		var skate_brand models.BrandStruct

		err = rows.Scan(
			&skate.ID,

			&model.ID,
			&model.Name,
			&model.Alias,

			&skate_brand.ID,
			&skate_brand.Name,
			&skate_brand.ShortName,
			&skate_brand.IsSteel,
			&skate_brand.IsSkate,
			&skate_brand.IsHolder,
		)
		if err != nil {
			return nil, err
		}

		skate.Model = model
		skate.Brand = skate_brand

		skatesArray = append(skatesArray, skate)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return skatesArray, nil
}

func getSkatesFromQuery(query *sql.Stmt) ([]models.SkateStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	skatesArray := make([]models.SkateStruct, 0)
	for rows.Next() {
		var skate models.SkateStruct

		var model models.ModelStruct

		var skate_brand models.BrandStruct

		err = rows.Scan(
			&skate.ID,

			&model.ID,
			&model.Name,
			&model.Alias,

			&skate_brand.ID,
			&skate_brand.Name,
			&skate_brand.ShortName,
			&skate_brand.IsSteel,
			&skate_brand.IsSkate,
			&skate_brand.IsHolder,
		)
		if err != nil {
			return nil, err
		}

		skate.Model = model
		skate.Brand = skate_brand

		skatesArray = append(skatesArray, skate)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return skatesArray, nil
}

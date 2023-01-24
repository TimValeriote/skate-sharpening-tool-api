package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

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

					skate_model_brand.id,
					skate_model_brand.name, 
					skate_model_brand.short_name, 
					skate_model_brand.is_skate, 
					skate_model_brand.is_steel, 
					skate_model_brand.is_holder,

					skate_brand.id,
					skate_brand.name, 
					skate_brand.short_name, 
					skate_brand.is_skate, 
					skate_brand.is_steel, 
					skate_brand.is_holder,

					skate_fit.id,
					skate_fit.name
			FROM skates
			INNER JOIN model AS skate_model ON skates.model_id = skate_model.id
			INNER JOIN brands AS skate_model_brand ON skate_model.brand_id = skate_model_brand.id
			INNER JOIN brands AS skate_brand ON skates.brand_id = skate_brand.id
			INNER JOIN fits AS skate_fit ON skates.fit_id = skate_fit.id`
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
		var skate_model_brand models.BrandStruct

		var skate_brand models.BrandStruct

		var fit models.FitStruct

		err = rows.Scan(
			&skate.ID,

			&model.ID,
			&model.Name,
			&model.Alias,

			&skate_model_brand.ID,
			&skate_model_brand.Name,
			&skate_model_brand.ShortName,
			&skate_model_brand.IsSteel,
			&skate_model_brand.IsSkate,
			&skate_model_brand.IsHolder,

			&skate_brand.ID,
			&skate_brand.Name,
			&skate_brand.ShortName,
			&skate_brand.IsSteel,
			&skate_brand.IsSkate,
			&skate_brand.IsHolder,

			&fit.ID,
			&fit.Name,
		)
		if err != nil {
			return nil, err
		}

		model.Brand = skate_model_brand
		skate.Model = model
		skate.Brand = skate_brand
		skate.Fit = fit

		skatesArray = append(skatesArray, skate)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return skatesArray, nil
}

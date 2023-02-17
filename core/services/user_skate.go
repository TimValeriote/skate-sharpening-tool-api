package core

import (
	"database/sql"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type userSkateStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func UserSkateStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *userSkateStore {
	return &userSkateStore{
		database: db,
		log:      log,
	}
}

func (store *userSkateStore) GetAllUserSkatesByUserID(userId int) ([]models.UserSkateStruct, error) {
	sql := `SELECT 
				user_skates.id,

				skate.id,

				skate_model.id, 
				skate_model.name, 
				skate_model.alias,

				skate_brand.id, 
				skate_brand.name, 
				skate_brand.short_name,
				skate_brand.is_skate,
				skate_brand.is_steel,
				skate_brand.is_holder,

				holder_brand.id,
				holder_brand.name,
				holder_brand.short_name,
				holder_brand.is_skate,
				holder_brand.is_steel,
				holder_brand.is_holder,

				user_skates.holder_size,
				user_skates.skate_size,

				user_skate_fit.id,
				user_skate_fit.name,

				lace_colour.id,
				lace_colour.colour,

				steel_brand.id,
				steel_brand.name,
				steel_brand.short_name,
				steel_brand.is_skate,
				steel_brand.is_steel,
				steel_brand.is_holder,

				user_skates.has_guards,
				user_skates.has_steel,

				guard_colour.id,
				guard_colour.colour,

				user_skates.preferred_radius
			FROM user_skates
			INNER JOIN users ON user_skates.user_id = users.id
			INNER JOIN skates as skate ON user_skates.skate_id = skate.id
			INNER JOIN model as skate_model ON skate.model_id = skate_model.id
			INNER JOIN brands as skate_brand ON skate.brand_id = skate_brand.id
			INNER JOIN brands as holder_brand ON user_skates.holder_brand_id = holder_brand.id
			INNER JOIN fits as user_skate_fit ON user_skates.fit_id = user_skate_fit.id
			INNER JOIN colour as lace_colour ON user_skates.lace_colour_id = lace_colour.id
			INNER JOIN brands as steel_brand ON user_skates.steel_id = steel_brand.id
			LEFT JOIN colour as guard_colour ON user_skates.guard_colour_id = guard_colour.id
			WHERE user_id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::GetAllUserSkatesByUserID - Failed to prepare GetAllUserSkatesByUserID SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	skates := make([]models.UserSkateStruct, 0)
	for rows.Next() {
		var userSkate models.UserSkateStruct

		var skate models.SkateStruct
		var skateBrand models.BrandStruct
		var skateModel models.ModelStruct

		var holder models.BrandStruct

		var laceColour models.ColourStruct

		var skate_steel models.BrandStruct

		var guardColour models.ColourStruct

		var fit models.FitStruct

		err = rows.Scan(
			&userSkate.ID,

			&skate.ID,

			&skateModel.ID,
			&skateModel.Name,
			&skateModel.Alias,

			&skateBrand.ID,
			&skateBrand.Name,
			&skateBrand.ShortName,
			&skateBrand.IsSkate,
			&skateBrand.IsSteel,
			&skateBrand.IsHolder,

			&holder.ID,
			&holder.Name,
			&holder.ShortName,
			&holder.IsSkate,
			&holder.IsSteel,
			&holder.IsHolder,

			&userSkate.HolderSize,
			&userSkate.SkateSize,

			&fit.ID,
			&fit.Name,

			&laceColour.ID,
			&laceColour.Colour,

			&skate_steel.ID,
			&skate_steel.Name,
			&skate_steel.ShortName,
			&skate_steel.IsSkate,
			&skate_steel.IsSteel,
			&skate_steel.IsHolder,

			&userSkate.HasGuards,
			&userSkate.HasSteel,

			&guardColour.ID,
			&guardColour.Colour,

			&userSkate.PreferredRadius,
		)
		if err != nil {
			return nil, err
		}

		skate.Model = skateModel
		skate.Brand = skateBrand
		userSkate.Fit = fit

		userSkate.Skate = skate
		userSkate.Holder = holder
		userSkate.LaceColour = laceColour
		userSkate.Steel = skate_steel
		userSkate.GuardColour = guardColour

		skates = append(skates, userSkate)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return skates, nil
}

func (store *userSkateStore) GetUserSkateByUserIdAndUserSkateId(userId int, userSkateId int) (models.UserSkateStruct, error) {
	sql := `SELECT 
				user_skates.id,

				skate.id,

				skate_model.id, 
				skate_model.name, 
				skate_model.alias,

				skate_brand.id, 
				skate_brand.name, 
				skate_brand.short_name,
				skate_brand.is_skate,
				skate_brand.is_steel,
				skate_brand.is_holder,

				holder_brand.id,
				holder_brand.name,
				holder_brand.short_name,
				holder_brand.is_skate,
				holder_brand.is_steel,
				holder_brand.is_holder,

				user_skates.holder_size,
				user_skates.skate_size,

				user_skate_fit.id,
				user_skate_fit.name,

				lace_colour.id,
				lace_colour.colour,

				steel_brand.id,
				steel_brand.name,
				steel_brand.short_name,
				steel_brand.is_skate,
				steel_brand.is_steel,
				steel_brand.is_holder,

				user_skates.has_guards,
				user_skates.has_steel,

				guard_colour.id,
				guard_colour.colour,

				user_skates.preferred_radius
			FROM user_skates
			INNER JOIN users ON user_skates.user_id = users.id
			INNER JOIN skates as skate ON user_skates.skate_id = skate.id
			INNER JOIN model as skate_model ON skate.model_id = skate_model.id
			INNER JOIN brands as skate_brand ON skate.brand_id = skate_brand.id
			INNER JOIN brands as holder_brand ON user_skates.holder_brand_id = holder_brand.id
			INNER JOIN fits as user_skate_fit ON user_skates.fit_id = user_skate_fit.id
			INNER JOIN colour as lace_colour ON user_skates.lace_colour_id = lace_colour.id
			INNER JOIN brands as steel_brand ON user_skates.steel_id = steel_brand.id
			LEFT JOIN colour as guard_colour ON user_skates.guard_colour_id = guard_colour.id
			WHERE user_skates.user_id = ? AND user_skates.id = ?`
	var userSkate models.UserSkateStruct
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::GetUserSkateByUserIdAndUserSkateId - Failed to prepare GetUserSkateByUserIdAndUserSkateId SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return userSkate, err
	}
	defer query.Close()

	rows, err := query.Query(userId, userSkateId)
	if err != nil {
		return userSkate, err
	}
	defer rows.Close()

	for rows.Next() {
		var skate models.SkateStruct
		var skateBrand models.BrandStruct
		var skateModel models.ModelStruct

		var holder models.BrandStruct

		var laceColour models.ColourStruct

		var skate_steel models.BrandStruct

		var guardColour models.ColourStruct

		var fit models.FitStruct

		err = rows.Scan(
			&userSkate.ID,

			&skate.ID,

			&skateModel.ID,
			&skateModel.Name,
			&skateModel.Alias,

			&skateBrand.ID,
			&skateBrand.Name,
			&skateBrand.ShortName,
			&skateBrand.IsSkate,
			&skateBrand.IsSteel,
			&skateBrand.IsHolder,

			&holder.ID,
			&holder.Name,
			&holder.ShortName,
			&holder.IsSkate,
			&holder.IsSteel,
			&holder.IsHolder,

			&userSkate.HolderSize,
			&userSkate.SkateSize,

			&fit.ID,
			&fit.Name,

			&laceColour.ID,
			&laceColour.Colour,

			&skate_steel.ID,
			&skate_steel.Name,
			&skate_steel.ShortName,
			&skate_steel.IsSkate,
			&skate_steel.IsSteel,
			&skate_steel.IsHolder,

			&userSkate.HasGuards,
			&userSkate.HasSteel,

			&guardColour.ID,
			&guardColour.Colour,

			&userSkate.PreferredRadius,
		)
		if err != nil {
			return userSkate, err
		}

		skate.Model = skateModel
		skate.Brand = skateBrand
		userSkate.Fit = fit

		userSkate.Skate = skate
		userSkate.Holder = holder
		userSkate.LaceColour = laceColour
		userSkate.Steel = skate_steel
		userSkate.GuardColour = guardColour
	}

	if rows.Err() != nil {
		return userSkate, err
	}

	return userSkate, nil
}

func (store *userSkateStore) CreateUserSkate(userSkateResponse *models.CreateUserSkateStruct) (userSkateId int, err error) {
	sql := `INSERT INTO user_skates (
		user_id,
		skate_id,
		holder_brand_id,
		holder_size,
		skate_size,
		lace_colour_id,
		has_steel,
		steel_id,
		has_guards,
		guard_colour_id,
		fit_id,
		preferred_radius
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::CreateUserSkate - Failed to prepare CreateUserSkate SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	res, err := sqlStmt.Exec(
		userSkateResponse.UserID,
		userSkateResponse.SkateID,
		userSkateResponse.HolderBrandID,
		userSkateResponse.HolderSize,
		userSkateResponse.SkateSize,
		userSkateResponse.LaceColourID,
		userSkateResponse.HasSteel,
		userSkateResponse.SteelID,
		userSkateResponse.HasGuards,
		userSkateResponse.GuardColourID,
		userSkateResponse.FitID,
		userSkateResponse.PreferredRadius,
	)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::CreateUserSkate - Failed to execute CreateUserSkate SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::CreateUserSkate - Failed to get the last inserted id",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userSkateId = int(insertedId)

	return
}

func (store *userSkateStore) UpdateUserSkate(userSkate models.UpdateUserSkateStruct) (result sql.Result, err error) {
	sql := `UPDATE user_skates SET 
					skate_id = ?, 
					holder_brand_id = ?,
					holder_size = ?,
					skate_size = ?,
					lace_colour_id = ?,
					has_steel = ?,
					steel_id = ?,
					has_guards = ?,
					guard_colour_id = ?,
					fit_id = ?,
					preferred_radius = ?
			WHERE id = ?`

	result, err = store.database.Tx.Exec(sql,
		userSkate.SkateID,
		userSkate.HolderBrandID,
		userSkate.HolderSize,
		userSkate.SkateSize,
		userSkate.LaceColourID,
		userSkate.HasSteel,
		userSkate.SteelID,
		userSkate.HasGuards,
		userSkate.GuardColourID,
		userSkate.FitID,
		userSkate.PreferredRadius,
		userSkate.UserSkateID,
	)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::UpdateUserSkate - Failed to execute UpdateUserSkate SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	return
}

func (store *userSkateStore) DeleteUserSkate(userId int, userSkateId int) (result sql.Result, err error) {
	sql := `DELETE FROM user_skates WHERE user_id = ? AND id = ?`

	result, err = store.database.Tx.Exec(sql, userId, userSkateId)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::DeleteUserSkate - Failed to execute DeleteUserSkate SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	return
}

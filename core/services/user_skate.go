package core

import (
	//"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

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

/*func (store *userSkateStore) GetPersonByID(personID int) ([]models.PeopleStruct, error) {
	sql := `SELECT id, first_name, last_name FROM people WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::GetPersonByID - Failed to prepare GetPersonByID SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]models.PeopleStruct, 0)
	for rows.Next() {
		var person models.PeopleStruct
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return people, nil
}

func (store *userSkateStore) CreatePerson(personResponse *models.PeopleStruct) (err error) {
	sql := `INSERT INTO people(first_name, last_name) VALUES (?,?)`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::CreatePerson - Failed to prepare CreatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	_, err = sqlStmt.Exec(personResponse.FirstName, personResponse.LastName)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::CreatePerson - Failed to execute CreatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	return nil
}

func (store *userSkateStore) UpdatePerson(person *models.PeopleStruct) (err error) {
	sql := `UPDATE people SET first_name = ?, last_name = ? WHERE id = ?`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::UpdatePerson - Failed to prepare UpdatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	_, err = sqlStmt.Exec(person.FirstName, person.LastName, person.ID)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userSkateStore::UpdatePerson - Failed to execute UpdatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	return nil
}

func getPeopleFromQuery(query *sql.Stmt) ([]models.PeopleStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]models.PeopleStruct, 0)
	for rows.Next() {
		var person models.PeopleStruct
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return people, nil
}*/

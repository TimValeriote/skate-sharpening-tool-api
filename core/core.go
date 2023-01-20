package core

import (
	"context"
	"database/sql"

	"phl-skate-sharpening-api/core/implementation"
	"phl-skate-sharpening-api/core/models"
)

func CreateCore(db *sql.DB) (*models.Core, error) {
	return CreateBACoreContext(context.Background(), db)
}

func CreateBACoreContext(ctx context.Context, db *sql.DB) (*models.Core, error) {
	var coreDB models.CoreDatabase

	core := &models.Core{
		Ctx: ctx,
	}

	coreDB.DB = db
	core.Database = &coreDB

	core.PeopleService = &implementation.PeopleStore{core}

	return core, nil

}

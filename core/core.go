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

	core.UserService = &implementation.UserStore{core}
	core.ColourService = &implementation.ColourStore{core}
	core.UserSkateService = &implementation.UserSkateStore{core}
	core.FitService = &implementation.FitStore{core}
	core.BrandService = &implementation.BrandStore{core}
	core.ModelService = &implementation.ModelStore{core}
	core.SkateService = &implementation.SkateStore{core}
	core.StoreService = &implementation.StoreStore{core}
	core.SharpeningService = &implementation.SharpeningStore{core}

	return core, nil
}

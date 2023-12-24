package postgres_dao

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type AppRegionsDAO struct {
	*sqlx.DB
}

func NewAppRegionsDAO(db *sqlx.DB) *AppRegionsDAO {
	return &AppRegionsDAO{db}
}

func (dao *AppRegionsDAO) InsertAppRegion(app_region *models.AppRegionMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO app_regions (app_id, region_id, created_time, updated_time) values (:app_id, :region_id, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, app_region)
	if err != nil {
		return err
	}

	return nil
}

func (dao *AppRegionsDAO) CheckAppRegionExist(app_region *models.AppRegionMD) bool {
	count := int8(0)
	fmt.Println(app_region.AppID, app_region.RegionID)
	query := "SELECT COUNT(1) from app_regions where app_id = $1 AND region_id = $2"
	err := dao.DB.Get(&count, query, app_region.AppID, app_region.RegionID)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionsDAO::CheckAppRegionExist  - Error: %+v", err)
		return base.Int8ToBool(count)
	}
	return base.Int8ToBool(count)
}

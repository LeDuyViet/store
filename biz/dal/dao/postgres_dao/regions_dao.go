package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type RegionsDAO struct {
	*sqlx.DB
}

func NewRegionsDAO(db *sqlx.DB) *RegionsDAO {
	return &RegionsDAO{db}
}

func (dao *RegionsDAO) InsertRegions(region *models.RegionMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO regions (id, name, code, created_time, updated_time) values (:id, :name, :code, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, region)
	if err != nil {
		return err
	}

	return nil
}

func (dao *RegionsDAO) GetRegionById(id string) *models.RegionMD {
	query := "SELECT id, name, code, created_time, updated_time from regions where id = $1"
	do, err := queryDataParser[models.RegionMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionsDAO::GetRegions - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *RegionsDAO) CheckRegionExist(condition string) bool {

	count := int8(0)
	query := "SELECT COUNT(1) from regions where id = $1 OR name = $1 OR code = $1"
	err := dao.DB.Get(&count, query, condition)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionsDAO::CheckRegionExist - Error: %+v", err)
		return base.Int8ToBool(count)
	}

	return base.Int8ToBool(count)
}

func (dao *RegionsDAO) GetRegions() []*models.RegionMD {
	query := "SELECT id, name, code, created_time, updated_time from regions"
	do, err := queryListDataParser[models.RegionMD](dao.DB, query, nil)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionsDAO::GetRegions - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *RegionsDAO) UpdateRegion(region *models.RegionMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE regions SET name = :name, code = :code, updated_time = :updated_time where id = :id"
	_, err := dao.NamedExecContext(ctx, query, region)
	if err != nil {
		v_log.V(1).Errorf("RegionsDAO::UpdateRegion - Error: %+v", err)
		return err
	}

	return nil
}

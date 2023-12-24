package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
)

type DataMigratesDAO struct {
	*sqlx.DB
}

func NewDataMigratesDAO(db *sqlx.DB) *DataMigratesDAO {
	return &DataMigratesDAO{db}
}

func (dao *DataMigratesDAO) InsertDataMigrate(dataMigrate *models.DataMigratesMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO data_migrates (store_id, data_id, data_type, migration_id, is_sync, created_time, updated_time) values (:store_id, :data_id, :data_type, :migration_id, :is_sync, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, dataMigrate)
	if err != nil {
		return err
	}

	return nil
}

func (dao *DataMigratesDAO) GetByMigrateData(dataMigrate *models.DataMigratesMD) *models.DataMigratesMD {
	query := "SELECT migration_id from data_migrates where store_id = $1 and data_id = $2 and data_type = $3"
	do, err := queryDataParser[models.DataMigratesMD](dao.DB, query, nil, dataMigrate.StoreID, dataMigrate.DataID, dataMigrate.DataType)
	if err != nil {
		return nil
	}

	do.DataID = dataMigrate.DataID
	do.StoreID = dataMigrate.StoreID
	do.DataType = dataMigrate.DataType

	return do
}

func (dao *DataMigratesDAO) GetByID(id string) *models.DataMigratesMD {
	query := "SELECT store_id, data_id, data_type, migration_id, is_sync, created_time, updated_time from data_migrates where migration_id = $1"
	do, err := queryDataParser[models.DataMigratesMD](dao.DB, query, nil, id)
	if err != nil {
		return nil
	}

	return do
}

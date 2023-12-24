package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type StoresDAO struct {
	*sqlx.DB
}

func NewStoresDAO(db *sqlx.DB) *StoresDAO {
	return &StoresDAO{db}
}

func (dao *StoresDAO) InsertStore(store *models.StoreMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO stores (id, name, address, is_active, access_key, created_time, updated_time) values (:id, :name, :address, :is_active, :access_key, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, store)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::InsertStore - Error: %+v", err)
		return err
	}

	return nil
}

func (dao *StoresDAO) GetStore(id string) *models.StoreMD {
	query := "SELECT id, name, address, is_active, access_key, created_time, updated_time from stores where id = $1"
	do, err := queryDataParser[models.StoreMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::GetStore - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *StoresDAO) GetByPackageName(packageName string) []*models.StoreMD {
	query := "SELECT id, name, address, is_active, access_key, created_time, updated_time from stores where package_name = $1"
	res, err := queryListDataParser[models.StoreMD](dao.DB, query, nil, packageName)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::GetByPackageName - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *StoresDAO) GetAll() []*models.StoreMD {
	query := "SELECT id, name, address, is_active, access_key, created_time, updated_time from stores"
	res, err := queryListDataParser[models.StoreMD](dao.DB, query, nil)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::GetStores - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *StoresDAO) GetStores(offset, limit int) []*models.StoreMD {
	query := "SELECT id, name, address, is_active, access_key, created_time, updated_time from stores LIMIT $1 OFFSET $2"
	do, err := queryListDataParser[models.StoreMD](dao.DB, query, nil, limit, offset)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::GetStores - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *StoresDAO) UpdateStore(store *models.StoreMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE stores SET name = :name, address = :address, is_active = :is_active, access_key = :access_key, updated_time = :updated_time where id = :id"
	_, err := dao.NamedExecContext(ctx, query, store)
	if err != nil {
		v_log.V(1).Errorf("StoresDAO::UpdateApp - Error: %+v", err)
		return err
	}

	return nil
}

func (dao *StoresDAO) DeleteStore(id string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "DELETE FROM stores where id = $1"
	_, err := dao.QueryxContext(ctx, query, id)
	if err != nil {
		v_log.V(1).Errorf("StoresDAO::DeleteStore - Error: %+v", err)
		return err
	}

	return nil
}

func (dao *StoresDAO) GetStoresByAddress(address string) *models.StoreMD {
	query := "SELECT id, name, address, is_active, access_key, created_time, updated_time from stores where address = $1"
	res, err := queryDataParser[models.StoreMD](dao.DB, query, nil, address)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("StoresDAO::GetStoresByAddress - Error: %+v", err)

		return nil
	}

	return res
}

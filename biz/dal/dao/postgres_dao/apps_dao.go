package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type AppsDAO struct {
	*sqlx.DB
}

func NewAppsDAO(db *sqlx.DB) *AppsDAO {
	return &AppsDAO{db}
}

func (dao *AppsDAO) InsertApp(app *models.AppMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO apps (id, store_id, name, package_name, photo, created_time, updated_time) values (:id, :store_id, :name, :package_name, :photo, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, app)
	if err != nil {
		return err
	}

	return nil
}

func (dao *AppsDAO) GetApp(id string) *models.AppMD {
	query := "SELECT id, store_id, name, package_name, photo, created_time, updated_time from apps where id = $1"
	do, err := queryDataParser[models.AppMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::GetApp - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *AppsDAO) GetAllAppsByStoreID(storeID string) []*models.AppMD {
	query := "SELECT id, store_id, name, package_name, photo, created_time, updated_time from apps where store_id = $1"
	res, err := queryListDataParser[models.AppMD](dao.DB, query, nil, storeID)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::GetAllAppsByStoreID - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *AppsDAO) GetByStore(storeID string, offset, limit int) []*models.AppMD {
	query := "SELECT id, store_id, name, package_name, photo, created_time, updated_time from apps where store_id = $1 limit $2 offset $3"
	res, err := queryListDataParser[models.AppMD](dao.DB, query, nil, storeID, limit, offset)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::GetAppsByStoreID - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *AppsDAO) GetAll(offset, limit int) []*models.AppMD {
	query := "SELECT id, store_id, name, package_name, photo, created_time, updated_time from apps limit $1 offset $2"
	res, err := queryListDataParser[models.AppMD](dao.DB, query, nil, limit, offset)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::GetAll - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *AppsDAO) Update(app *models.AppMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE apps SET name = :name, package_name = :package_name, photo = :photo, updated_time = :updated_time WHERE id = :id"
	_, err := dao.NamedExecContext(ctx, query, app)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::Update - Error: %+v", err)

		return err
	}

	return nil
}

// get by package name
func (dao *AppsDAO) GetByPackageName(packageName string) []*models.AppMD {
	query := "SELECT id, store_id, name, package_name, photo, created_time, updated_time from apps where package_name = $1"
	res, err := queryListDataParser[models.AppMD](dao.DB, query, nil, packageName)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("AppsDAO::GetByPackageName - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *AppsDAO) Delete(id string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "DELETE FROM apps where id = $1"
	_, err := dao.QueryxContext(ctx, query, id)
	if err != nil {
		v_log.V(1).Errorf("AppsDAO::Delete - Error: %+v", err)
		return err
	}

	return nil
}

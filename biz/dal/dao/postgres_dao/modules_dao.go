package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type ModulesDAO struct {
	*sqlx.DB
}

func NewModulesDAO(db *sqlx.DB) *ModulesDAO {
	return &ModulesDAO{db}
}

func (dao *ModulesDAO) InsertModule(app *models.ModuleMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO modules (id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time) VALUES (:id, :app_id, :priority, :name, :thumbnail, :icon, :status, :is_pro, :is_new, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, app)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ModulesDAO) GetModule(id string) *models.ModuleMD {
	query := "SELECT id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM modules WHERE id = $1"
	do, err := queryDataParser[models.ModuleMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::GetModule - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *ModulesDAO) GetById(app_id string, per_page, page int) ([]*models.ModuleMD, int) {
	do := []*models.ModuleMD{}
	var err error
	if per_page != -1 {
		query := "SELECT id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM modules WHERE app_id = $1 and status = true ORDER By created_time ASC LIMIT $2 OFFSET $3 "
		offset := per_page * (page - 1)
		do, err = queryListDataParser[models.ModuleMD](dao.DB, query, nil, app_id, per_page, offset)

	} else {
		query := "SELECT id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM modules WHERE app_id = $1 and status = true ORDER By created_time ASC"
		do, err = queryListDataParser[models.ModuleMD](dao.DB, query, nil, app_id)
	}
	total := countTotal(dao.DB, "modules", "app_id = $1 and status = true", app_id)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::GetModule - Error: %+v", err)

		return nil, 0
	}

	return do, total
}

func (dao *ModulesDAO) UpdatePriority(id string, priority int32) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()
	query := "UPDATE modules SET priority = $1 WHERE id = $2"
	_, err := dao.QueryxContext(ctx, query, priority, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::UpdatePriority - Error: %+v - id: %s", err, id)
		return err
	}

	return nil
}

func (dao *ModulesDAO) Update(app *models.ModuleMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()
	query := "UPDATE modules SET app_id = :app_id, priority = :priority, name = :name, thumbnail = :thumbnail, icon = :icon, status = :status, is_pro = :is_pro, is_new = :is_new, updated_time = :updated_time WHERE id = :id"
	_, err := dao.NamedExecContext(ctx, query, app)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::Update - Error: %+v - id: %s", err, app.ID)
		return err
	}

	return nil
}

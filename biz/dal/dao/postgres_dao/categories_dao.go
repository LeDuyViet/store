package postgres_dao

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type CategoriesDAO struct {
	*sqlx.DB
}

func NewCategoriesDAO(db *sqlx.DB) *CategoriesDAO {
	return &CategoriesDAO{db}
}

func (dao *CategoriesDAO) InsertCategory(category *models.CategoryMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO categories (id, parent_id, level, module_id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time) VALUES (:id, :parent_id, :level, :module_id, :app_id, :priority, :name, :thumbnail, :icon, :status, :is_pro, :is_new, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, category)
	if err != nil {
		return err
	}

	return nil
}

func (dao *CategoriesDAO) GetCategory(id string) *models.CategoryMD {
	query := "SELECT * FROM categories WHERE id = $1"
	do, err := queryDataParser[models.CategoryMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::GetCategory - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *CategoriesDAO) GetCategories(offset, limit int) []*models.CategoryMD {
	query := "SELECT * FROM categories LIMIT $1 OFFSET $2"
	res, err := queryListDataParser[models.CategoryMD](dao.DB, query, nil, limit, offset)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::GetCategories - Error: %+v", err)

		return nil
	}

	return res
}
func (dao *CategoriesDAO) GetCategorybyModuleId(module_id string, per_page, page int) ([]*models.CategoryMD, int) {
	do := []*models.CategoryMD{}
	var err error
	if per_page != -1 {
		query := "SELECT id, module_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM categories  WHERE module_id = $1 and status = true ORDER By created_time ASC LIMIT $2 OFFSET $3"
		offset := per_page * (page - 1)
		do, err = queryListDataParser[models.CategoryMD](dao.DB, query, nil, module_id, per_page, offset)

	} else {
		query := "SELECT id, module_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM categories  WHERE module_id = $1 and status = true ORDER By created_time ASC"
		do, err = queryListDataParser[models.CategoryMD](dao.DB, query, nil, module_id)
	}
	total := countTotal(dao.DB, "categories", "module_id = $1 and status = true", module_id)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::GetModule - Error: %+v", err)

		return nil, 0
	}

	return do, total
}

func (dao *CategoriesDAO) GetCategorybyModuleIdForAdmin(module_id string, per_page, page int) ([]*models.CategoryMD, int) {
	do := []*models.CategoryMD{}
	var err error
	if per_page != -1 {
		query := "SELECT id, module_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM categories  WHERE module_id = $1 ORDER By created_time ASC LIMIT $2 OFFSET $3"
		offset := per_page * (page - 1)
		do, err = queryListDataParser[models.CategoryMD](dao.DB, query, nil, module_id, per_page, offset)

	} else {
		query := "SELECT id, module_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM categories  WHERE module_id = $1 ORDER By created_time ASC"
		do, err = queryListDataParser[models.CategoryMD](dao.DB, query, nil, module_id)
	}
	total := countTotal(dao.DB, "categories", "module_id = $1 and status = true", module_id)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("ModulesDAO::GetModule - Error: %+v", err)

		return nil, 0
	}

	return do, total
}

func (dao *CategoriesDAO) UpdatePriority(id string, priority int32) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()
	query := "UPDATE categories SET priority = $1 WHERE id = $2"
	_, err := dao.QueryxContext(ctx, query, priority, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::UpdatePriority - Error: %+v - id: %s", err, id)
		return err
	}

	return nil
}

func (dao *CategoriesDAO) CheckCategoryExist(id string) bool {

	count := int8(0)
	query := "SELECT COUNT(1) from categories where id = $1"
	err := dao.DB.Get(&count, query, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::CheckCategoriesExist - Error: %+v", err)
		return base.Int8ToBool(count)
	}

	return base.Int8ToBool(count)
}

func (dao *CategoriesDAO) GetByApp(appID string, offset, limit int, onlyRoot ...bool) []*do.CategoriesDO {
	var query string

	if len(onlyRoot) > 0 && onlyRoot[0] {
		query = "SELECT id, parent_id, priority, name, thumbnail, icon, status, is_pro, is_new FROM categories where app_id = $1 and id = parent_id offset $2 limit $3"
	} else {
		query = "SELECT id, parent_id, priority, name, thumbnail, icon, status, is_pro, is_new FROM categories where app_id = $1 offset $2 limit $3"
	}

	do, err := queryListDataParser[do.CategoriesDO](dao.DB, query, nil, appID, offset, limit)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::GetByApp - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *CategoriesDAO) GetByParent(parentID string) []*do.CategoriesDO {
	query := "SELECT id, parent_id, priority, name, thumbnail, icon, status, is_pro, is_new FROM categories where parent_id = $1 and id <> parent_id"
	do, err := queryListDataParser[do.CategoriesDO](dao.DB, query, nil, parentID)

	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::GetByParent - Error: %+v", err)
	}

	return do
}

func (dao *CategoriesDAO) GetAllByApp(appID string) []*models.CategoryMD {
	query := "SELECT id, parent_id, level, module_id, app_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM categories where app_id = $1"
	do, err := queryListDataParser[models.CategoryMD](dao.DB, query, nil, appID)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::GetAllByApp - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *CategoriesDAO) Update(category *models.CategoryMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE categories SET id = :id, parent_id = :parent_id, level = :level, module_id = :module_id, app_id = :app_id, priority = :priority, name = :name, thumbnail = :thumbnail, icon = :icon, status = :status, is_pro = :is_pro, is_new = :is_new, updated_time = :updated_time WHERE id = :id"
	_, err := dao.NamedExecContext(ctx, query, category)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::Update - Error: %+v", err)

		return err
	}

	return nil
}

func (dao *CategoriesDAO) Delete(id string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "DELETE FROM categories where id = $1"
	_, err := dao.QueryxContext(ctx, query, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::Delete - Error: %+v", err)

		return err
	}

	return nil
}

func (dao *CategoriesDAO) DeleteMultiple(ids []string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	tx := dao.MustBeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	for _, id := range ids {
		_, err := tx.Exec("DELETE FROM categories WHERE id = $1", id)
		if err != nil {
			v_log.V(1).WithError(err).Errorf("CategoriesDAO::DeleteMultiple - Error: %+v", err)
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CategoriesDAO::DeleteMultiple - Error: %+v", err)
		return err
	}

	return nil
}

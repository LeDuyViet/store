package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type ItemsDAO struct {
	*sqlx.DB
}

func NewItemsDAO(db *sqlx.DB) *ItemsDAO {
	return &ItemsDAO{db}
}

func (dao *ItemsDAO) InsertItem(item *models.ItemMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO items (id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time) values (:id, :category_id, :priority, :name, :thumbnail, :icon, :status, :is_pro, :is_new, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, item)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ItemsDAO) GetItem(id string) *models.ItemMD {
	query := "SELECT id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time from items where id = $1"
	do, err := queryDataParser[models.ItemMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ItemsDAO::GetItem - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *ItemsDAO) GetItemsbyCategoryId(category_id string, per_page, page int) ([]*models.ItemMD, int) {
	do := []*models.ItemMD{}
	var err error
	if per_page != -1 {
		query := "SELECT id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM items  WHERE category_id = $1 and status = true ORDER By priority ASC LIMIT $2 OFFSET $3"
		offset := per_page * (page - 1)
		do, err = queryListDataParser[models.ItemMD](dao.DB, query, nil, category_id, per_page, offset)

	} else {
		query := "SELECT id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM items  WHERE category_id = $1 and status = true ORDER By priority ASC"
		do, err = queryListDataParser[models.ItemMD](dao.DB, query, nil, category_id)
	}
	total := countTotal(dao.DB, "items", "category_id = $1 and status = true", category_id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ItemsDAO::GetItemsbyCategoryId - Error: %+v", err)

		return nil, 0
	}

	return do, total
}

func (dao *ItemsDAO) GetItemsbyCategoryIdForAdmin(category_id string, per_page, page int) ([]*models.ItemMD, int) {
	do := []*models.ItemMD{}
	var err error
	if per_page != -1 {
		query := "SELECT id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM items  WHERE category_id = $1 ORDER By priority ASC LIMIT $2 OFFSET $3"
		offset := per_page * (page - 1)
		do, err = queryListDataParser[models.ItemMD](dao.DB, query, nil, category_id, per_page, offset)

	} else {
		query := "SELECT id, category_id, priority, name, thumbnail, icon, status, is_pro, is_new, created_time, updated_time FROM items  WHERE category_id = $1 ORDER By priority ASC"
		do, err = queryListDataParser[models.ItemMD](dao.DB, query, nil, category_id)
	}
	total := countTotal(dao.DB, "items", "category_id = $1 ", category_id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("ItemsDAO::GetItemsbyCategoryId - Error: %+v", err)

		return nil, 0
	}

	return do, total
}

func (dao *ItemsDAO) UpdatePriority(id string, priority int) error {
	query := "UPDATE items SET priority = $1 WHERE id = $2"
	_, err := queryDataParser[models.ModuleMD](dao.DB, query, nil, priority, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ItemsDAO) CheckItemExist(condition string) bool {

	count := int8(0)
	query := "SELECT COUNT(1) from items where id = $1"
	err := dao.DB.Get(&count, query, condition)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("RegionsDAO::CheckItemsExist - Error: %+v", err)
		return base.Int8ToBool(count)
	}

	return base.Int8ToBool(count)
}

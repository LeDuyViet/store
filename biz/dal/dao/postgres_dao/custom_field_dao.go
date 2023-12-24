package postgres_dao

import (
	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type CustomFieldsDAO struct {
	*sqlx.DB
}

func NewCustomFieldsDAO(db *sqlx.DB) *CustomFieldsDAO {
	return &CustomFieldsDAO{db}
}

func (dao *CustomFieldsDAO) InsertCustomField(customField *models.CustomFieldMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO custom_fields (id, name, app_id, type, is_module, is_category, is_item, created_time, updated_time) VALUES (:id, :name, :app_id, :type, :is_module, :is_category, :is_item, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, customField)
	if err != nil {
		return err
	}

	return nil
}

func (dao *CustomFieldsDAO) GetCustomField(id string) *models.CustomFieldMD {
	query := "SELECT * FROM custom_fields WHERE id = $1"
	do, err := queryDataParser[models.CustomFieldMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldsDAO::GetCustomField - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *CustomFieldsDAO) GetForModule(module_id, app_id string) []*do.CustomFieldTable {
	query := `select 
				cft.id, 
				custom_field_id, 
				name, 
				type, 
				is_active, 
				value 
  			from 
				custom_fields cf 
			inner join custom_field_tables cft on cf.id = cft.custom_field_id 
				and cft.custom_field_tableable_id = $1 
				and cft.custom_field_tableable_type = $2 
  			where 
				cf.app_id = $3`
	do, err := queryListDataParser[do.CustomFieldTable](dao.DB, query, nil, module_id, models.KTypeModule, app_id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldsDAO::GetCustomFieldForModule - Error: %+v", err)
		return nil
	}
	return do
}

func (dao *CustomFieldsDAO) GetCustomFieldForCategory(category_id, module_id string) []*do.CustomFieldTable {
	query := `select 
				cft.id,
				custom_field_id,
				value,
				is_active,
				cf.name,
				type
			from
				custom_fields cf
			inner join modules m on
 				m.id = $1
			inner join custom_field_tables cft on
				cf.id = cft.custom_field_id
				and cft.custom_field_tableable_type = $2
				and cft.custom_field_tableable_id = $3
			where
				cf.app_id = m.app_id`
	do, err := queryListDataParser[do.CustomFieldTable](dao.DB, query, nil, module_id, models.KTypeCategory, category_id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldsDAO::GetCustomFieldForModule - Error: %+v", err)
		return nil
	}
	return do
}

func (dao *CustomFieldsDAO) GetCustomFieldForItem(item_id, category_id string) []*do.CustomFieldTable {
	query := `select 
				cft.id,
				custom_field_id,
				value,
				is_active,
				cf.name,
				type
			from
				custom_fields cf
			inner join categories c on
				c.id = $1
			inner join modules m on
			 	m.id = c.module_id
			inner join custom_field_tables cft on
				cf.id = cft.custom_field_id
				and cft.custom_field_tableable_type = $2
				and cft.custom_field_tableable_id = $3
			where
				cf.app_id = m.app_id`
	do, err := queryListDataParser[do.CustomFieldTable](dao.DB, query, nil, category_id, models.KTypeItem, item_id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldsDAO::GetCustomFieldForItem - Error: %+v", err)
		return nil
	}
	return do
}

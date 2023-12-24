package postgres_dao

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

type CustomFieldDataDAO struct {
	*sqlx.DB
}

func NewCustomFieldDataDAO(db *sqlx.DB) *CustomFieldDataDAO {
	return &CustomFieldDataDAO{db}
}

func (dao *CustomFieldDataDAO) Insert(customField *models.CustomFieldDataMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO custom_field_tables (id, custom_field_id, value, custom_field_tableable_id, custom_field_tableable_type, is_active, created_time, updated_time) VALUES (:id, :custom_field_id, :value, :custom_field_tableable_id, :custom_field_tableable_type, :is_active, :created_time, :updated_time)"
	_, err := dao.NamedExecContext(ctx, query, customField)
	if err != nil {
		return err
	}

	return nil
}

func (dao *CustomFieldDataDAO) InsertMany(customFields []*models.CustomFieldDataMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "INSERT INTO custom_field_tables (id, custom_field_id, value, custom_field_tableable_id, custom_field_tableable_type, is_active, created_time, updated_time) VALUES (:id, :custom_field_id, :value, :custom_field_tableable_id, :custom_field_tableable_type, :is_active, :created_time, :updated_time)"

	tx := dao.MustBeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	_, err := tx.NamedExec(query, customFields)
	if err != nil {
		tx.Rollback()
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::InsertCustomFieldTable - Error: %+v", err)
	}

	err = tx.Commit()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::InsertCustomFieldTable - Error: %+v", err)
		return err
	}

	return nil
}

func (dao *CustomFieldDataDAO) UpdateMany(customFields []*models.CustomFieldDataMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE custom_field_tables SET custom_field_id = :custom_field_id, value = :value, custom_field_tableable_id = :custom_field_tableable_id, custom_field_tableable_type = :custom_field_tableable_type, is_active = :is_active WHERE id = :id"

	tx := dao.MustBeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	_, err := tx.NamedExecContext(ctx, query, customFields)
	if err != nil {
		tx.Rollback()
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::UpdateCustomFieldTable - Error: %+v", err)
	}

	err = tx.Commit()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::UpdateCustomFieldTable - Error: %+v", err)
	}

	return nil
}

func (dao *CustomFieldDataDAO) Update(customField *models.CustomFieldDataMD) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "UPDATE custom_field_tables SET custom_field_id = :custom_field_id, value = :value, custom_field_tableable_id = :custom_field_tableable_id, custom_field_tableable_type = :custom_field_tableable_type, is_active = :is_active WHERE id = :id"
	_, err := dao.NamedExecContext(ctx, query, customField)
	if err != nil {
		return err
	}

	return nil
}

func (dao *CustomFieldDataDAO) Get(id string) *models.CustomFieldDataMD {
	query := "SELECT * FROM custom_field_tables WHERE id = $1"
	do, err := queryDataParser[models.CustomFieldDataMD](dao.DB, query, nil, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::GetCustomFieldTable - Error: %+v", err)

		return nil
	}

	return do
}

func (dao *CustomFieldDataDAO) GetByIDs(customFieldIDs []string) []*models.CustomFieldDataMD {
	query := "SELECT id, value FROM custom_field_tables WHERE id IN (?)"

	var results []*models.CustomFieldDataMD

	// ids := make([]string, len(customFieldIDs))
	// for i, id := range customFieldIDs {
	// 	ids[i] = id
	// }

	query, args, err := sqlx.In(query, customFieldIDs)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::GetCustomFieldTable - Error: %+v", err)
		return nil
	}

	query = dao.DB.Rebind(query)
	err = dao.DB.Select(&results, query, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::GetCustomFieldTable - Error: %+v", err)

		return nil
	}

	return results
}

// return results
func (dao *CustomFieldDataDAO) GetMany(customFieldsDO []*do.CustomFieldTable) []*models.CustomFieldDataMD {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()
	var results []*models.CustomFieldDataMD

	whereClause := []string{}
	for _, cf := range customFieldsDO {
		whereClause = append(whereClause, fmt.Sprintf("(custom_field_id = '%s' AND custom_field_tableable_id = '%s' AND custom_field_tableable_type = '%s')", cf.CustomFieldID, cf.CustomFieldTableableID, cf.CustomFieldTableableType))
	}

	// Build the final query
	query := fmt.Sprintf("SELECT id, custom_field_id, value, custom_field_tableable_id, custom_field_tableable_type FROM custom_field_tables WHERE %s", strings.Join(whereClause, " OR "))

	// Execute the query
	rows, err := dao.DB.QueryxContext(ctx, query)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::GetCustomFieldTable - Error: %+v", err)
		return nil
	}
	defer rows.Close()

	// Retrieve the results
	for rows.Next() {
		var model models.CustomFieldDataMD
		err := rows.StructScan(&model)
		if err != nil {
			v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::GetCustomFieldTable - Error: %+v", err)
			return nil
		}
		results = append(results, &model)
	}

	return results

}

func (dao *CustomFieldDataDAO) GetByTableable(customFieldTableableID, customFieldTableableType string) []*models.CustomFieldDataMD {
	query := "SELECT id, custom_field_id, value, custom_field_tableable_id, custom_field_tableable_type FROM custom_field_tables WHERE custom_field_tableable_id = $1 AND custom_field_tableable_type = $2"
	res, err := queryListDataParser[models.CustomFieldDataMD](dao.DB, query, nil, customFieldTableableID, customFieldTableableType)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("GetByCustomFieldTableableIDAndType::GetAppsByStoreID - Error: %+v", err)

		return nil
	}

	return res
}

func (dao *CustomFieldDataDAO) Delete(id string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	query := "DELETE FROM custom_field_tables WHERE id = $1"
	_, err := dao.QueryxContext(ctx, query, id)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::DeleteCustomFieldTable - Error: %+v", err)

		return err
	}

	return nil
}

func (dao *CustomFieldDataDAO) DeleteMultiple(ids []string) error {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	tx := dao.MustBeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	for _, id := range ids {
		_, err := tx.Exec("DELETE FROM custom_field_tables WHERE id = $1", id)
		if err != nil {
			v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::DeleteCustomFieldTable - Error: %+v", err)
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		v_log.V(1).WithError(err).Errorf("CustomFieldDataDAO::DeleteCustomFieldTable - Error: %+v", err)
		return err
	}

	return nil
}

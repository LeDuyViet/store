package custom_fields

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

var _ core.CustomFieldsCallback = (*CustomFieldsController)(nil)

func (ctrl *CustomFieldsController) GetCustomField(id string) *models.CustomFieldMD {
	return ctrl.dao.customFieldDAO.GetCustomField(id)
}

func (ctrl *CustomFieldsController) CreateCustomField(customField *models.CustomFieldMD) *v_proto.VolioRpcError {
	if err := ctrl.dao.customFieldDAO.InsertCustomField(customField); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *CustomFieldsController) GetForModule(moduleID, appID string) []*do.CustomFieldTable {
	return ctrl.dao.customFieldDAO.GetForModule(moduleID, appID)
}

func (ctrl *CustomFieldsController) GetForPublicModule(moduleID, appID string) []*do.StoreCustomFieldTableDO {
	customFieldTables := ctrl.dao.customFieldDAO.GetForModule(moduleID, appID)
	customFieldsDO := make([]*do.StoreCustomFieldTableDO, 0)

	for _, ctf := range customFieldTables {
		customFieldTableMigrate := ctrl.dataMigratesCallback.GetByID(ctf.ID)
		customFieldMigrate := ctrl.dataMigratesCallback.GetByID(ctf.CustomFieldID)

		customFieldsDO = append(customFieldsDO, &do.StoreCustomFieldTableDO{
			ID:               customFieldTableMigrate.DataID,
			CustomFieldID:    customFieldMigrate.DataID,
			CustomFieldValue: ctf.CustomFieldValue,
			IsActive:         base.BoolToInt8(ctf.IsActive),
			Name:             ctf.Name,
			Type:             ctf.Type,
		})
	}

	return customFieldsDO
}

func (ctrl *CustomFieldsController) GetForPublicCategory(category_id, module_id string) []*do.StoreCustomFieldTableDO {
	customFieldTables := ctrl.dao.customFieldDAO.GetCustomFieldForCategory(category_id, module_id)
	customFieldsDO := make([]*do.StoreCustomFieldTableDO, 0)

	for _, ctf := range customFieldTables {
		customFieldTableMigrate := ctrl.dataMigratesCallback.GetByID(ctf.ID)
		customFieldMigrate := ctrl.dataMigratesCallback.GetByID(ctf.CustomFieldID)

		customFieldsDO = append(customFieldsDO, &do.StoreCustomFieldTableDO{
			ID:               customFieldTableMigrate.DataID,
			CustomFieldID:    customFieldMigrate.DataID,
			CustomFieldValue: ctf.CustomFieldValue,
			IsActive:         base.BoolToInt8(ctf.IsActive),
			Name:             ctf.Name,
			Type:             ctf.Type,
		})
	}

	return customFieldsDO
}

func (ctrl *CustomFieldsController) GetForPublicItem(category_id, module_id string) []*do.StoreCustomFieldTableDO {
	customFieldTables := ctrl.dao.customFieldDAO.GetCustomFieldForItem(category_id, module_id)
	customFieldsDO := make([]*do.StoreCustomFieldTableDO, 0)

	for _, ctf := range customFieldTables {
		customFieldTableMigrate := ctrl.dataMigratesCallback.GetByID(ctf.ID)
		customFieldMigrate := ctrl.dataMigratesCallback.GetByID(ctf.CustomFieldID)

		customFieldsDO = append(customFieldsDO, &do.StoreCustomFieldTableDO{
			ID:               customFieldTableMigrate.DataID,
			CustomFieldID:    customFieldMigrate.DataID,
			CustomFieldValue: ctf.CustomFieldValue,
			IsActive:         base.BoolToInt8(ctf.IsActive),
			Name:             ctf.Name,
			Type:             ctf.Type,
		})
	}

	return customFieldsDO
}

package custom_fields

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type customFieldDAO struct {
	customFieldDAO       *postgres_dao.CustomFieldsDAO
	customFieldTablesDAO *postgres_dao.CustomFieldDataDAO
}

type CustomFieldsController struct {
	dao *customFieldDAO

	dataMigratesCallback core.DataMigratesCallback
}

func (ctrl *CustomFieldsController) InstallController() {
	ctrl.dao.customFieldDAO = dao.GetCustomFieldsDAO(dao.STORES_DB_MASTER)
	ctrl.dao.customFieldTablesDAO = dao.GetCustomFieldDataDAO(dao.STORES_DB_MASTER)
}

func (ctrl *CustomFieldsController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.DataMigratesCallback:
		ctrl.dataMigratesCallback = x
	}
}

func (ctrl *CustomFieldsController) AfterInstalledDone() {
	v_log.V(3).Infof("CustomFieldsController::AfterInstalledDone")
}

func init() {
	f_core.RegisterCoreController(&CustomFieldsController{dao: &customFieldDAO{}})
}

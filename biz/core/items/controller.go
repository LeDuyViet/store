package items

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

const basePath = "%s/api/v2/public/items"

type itemsDAO struct {
	itemsDAO *postgres_dao.ItemsDAO
}

type ItemsController struct {
	dao                  *itemsDAO
	dataMigratesCallback core.DataMigratesCallback
	customFieldCallback  core.CustomFieldsCallback
}

func (ctrl *ItemsController) InstallController() {
	ctrl.dao.itemsDAO = dao.GetItemsDAO(dao.STORES_DB_MASTER)
}

func (ctrl *ItemsController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.CustomFieldsCallback:
		ctrl.customFieldCallback = x
	case core.DataMigratesCallback:
		ctrl.dataMigratesCallback = x
	}
}

func (ctrl *ItemsController) AfterInstalledDone() {
}

func init() {
	f_core.RegisterCoreController(&ItemsController{dao: &itemsDAO{}})
}

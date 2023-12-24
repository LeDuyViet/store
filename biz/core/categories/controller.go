package categories

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/admin"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

const basePath = "%s/api/v2/public/categories"

type categoriesDAO struct {
	categoriesDAO *postgres_dao.CategoriesDAO
}

type CategoriesController struct {
	dao *categoriesDAO

	*admin.AdminController

	customFieldCallback  core.CustomFieldsCallback
	dataMigratesCallback core.DataMigratesCallback
}

func (ctrl *CategoriesController) InstallController() {
	ctrl.dao.categoriesDAO = dao.GetCategoriesDAO(dao.STORES_DB_MASTER)
}

func (ctrl *CategoriesController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.CustomFieldsCallback:
		ctrl.customFieldCallback = x
	case core.DataMigratesCallback:
		ctrl.dataMigratesCallback = x
	case *admin.AdminController:
		ctrl.AdminController = x
	}
}

func (ctrl *CategoriesController) AfterInstalledDone() {
	v_log.V(3).Infof("CategoriesController::AfterInstalledDone")
	ctrl.registerCustomValidorRule()
}

func init() {
	f_core.RegisterCoreController(&CategoriesController{dao: &categoriesDAO{}})
}

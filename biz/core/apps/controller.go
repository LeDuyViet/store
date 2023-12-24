package apps

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/admin"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type appsDAO struct {
	appsDAO *postgres_dao.AppsDAO
}

type AppsController struct {
	dao *appsDAO

	*admin.AdminController

	categoriesCallback core.CategoriesCallback
}

func (ctrl *AppsController) InstallController() {
	ctrl.dao.appsDAO = dao.GetAppsDAO(dao.STORES_DB_MASTER)
}

func (ctrl *AppsController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.CategoriesCallback:
		ctrl.categoriesCallback = x
	case *admin.AdminController:
		ctrl.AdminController = x
	}

}

func (ctrl *AppsController) AfterInstalledDone() {
	ctrl.registerCustomValidorRule()
}

func init() {
	f_core.RegisterCoreController(&AppsController{dao: &appsDAO{}})
}

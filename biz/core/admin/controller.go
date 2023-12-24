package admin

import (
	"sync"

	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type AdminController struct {
	adminControllers sync.Map

	customFieldsCallback core.CustomFieldsCallback
}

func (ctrl *AdminController) InstallController() {
}

func (ctrl *AdminController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.StoresCallback:
		ctrl.adminControllers.Store(models.KTypeStoreData, x)
	case core.AppsCallback:
		ctrl.adminControllers.Store(models.KTypeAppData, x)
	case core.CategoriesCallback:
		ctrl.adminControllers.Store(models.KTypeCategory, x)
	case core.CustomFieldsCallback:
		ctrl.customFieldsCallback = x
	}
}

func (ctrl *AdminController) AfterInstalledDone() {
}

func init() {
	f_core.RegisterCoreController(&AdminController{adminControllers: sync.Map{}})
}

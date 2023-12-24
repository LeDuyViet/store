package apps

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
)

func (ctrl *AppsController) GetCategories(appID string, offset, limit int) []*do.CategoriesDO {
	return ctrl.categoriesCallback.GetByApp(appID, offset, limit)
}

func (ctrl *AppsController) GetAllCategories(appID string) []*models.CategoryMD {
	return ctrl.categoriesCallback.GetAllByApp(appID)
}

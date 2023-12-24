package categories_handler

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/categories"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/data_migrates"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
)

type CategoriesAPI struct {
	*v_api.VolioAPI

	categoriesController   *categories.CategoriesController
	dataMigratesController *data_migrates.DataMigratesController
}

func NewCategoriesAPI(controllers []core.CoreController) *CategoriesAPI {
	impl := &CategoriesAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *categories.CategoriesController:
			impl.categoriesController = x
		case *data_migrates.DataMigratesController:
			impl.dataMigratesController = x
		}
	}

	return impl
}

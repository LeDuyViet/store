package items_handler

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/data_migrates"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/items"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
)

type ItemsAPI struct {
	*v_api.VolioAPI

	itemsController        *items.ItemsController
	dataMigratesController *data_migrates.DataMigratesController
}

func NewItemsAPI(controllers []core.CoreController) *ItemsAPI {
	impl := &ItemsAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *items.ItemsController:
			impl.itemsController = x
		case *data_migrates.DataMigratesController:
			impl.dataMigratesController = x
		}
	}

	return impl
}

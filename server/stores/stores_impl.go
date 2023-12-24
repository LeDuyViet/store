package stores_handler

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/data_migrates"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/stores"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
)

type StoresAPI struct {
	*v_api.VolioAPI

	storesController       *stores.StoresController
	dataMigratesController *data_migrates.DataMigratesController
}

func NewStoresAPI(controllers []core.CoreController) *StoresAPI {
	impl := &StoresAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *stores.StoresController:
			impl.storesController = x
		case *data_migrates.DataMigratesController:
			impl.dataMigratesController = x
		}
	}

	return impl
}

package regions_handler

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/regions"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
)

type RegionAPI struct {
	*v_api.VolioAPI

	regionsController *regions.RegionsController
}

func NewRegionAPI(controllers []core.CoreController) *RegionAPI {
	impl := &RegionAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *regions.RegionsController:
			impl.regionsController = x
		}
	}

	return impl
}

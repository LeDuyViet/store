package regions_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
)

func (api *RegionAPI) GetAllRegions(ctx *fiber.Ctx) error {

	result := api.regionsController.GetAllRegions()
	return v_api.WriteSuccess(ctx, result)

}

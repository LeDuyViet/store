package regions_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
)

func (api *RegionAPI) Get(ctx *fiber.Ctx) error {
	client := store.GetStoreClient()
	region := ctx.Query("region")
	if region == "" {
		return ctx.Redirect(client.PublicHost + "/myzgos/api/v2/public/modules?package_name=com.callscreen.colorphone.calltheme.callscreener.callerscreenapp")
	}

	res := client.Login()

	return v_api.WriteSuccess(ctx, res)
}

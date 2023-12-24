package regions_handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *RegionAPI) UpdateRegions(ctx *fiber.Ctx) error {
	payload := models.RegionMD{}

	if err := ctx.BodyParser(&payload); err != nil {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	}

	if payload.ID == "" {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "id is Required!"))
	}

	if payload.Name == "" || payload.Code == "" {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "name or code is empty!"))
	}

	// validate country code
	url := fmt.Sprintf("https://restcountries.com/v3/alpha/%s", payload.Code)
	resp, _ := http.Get(url)
	if resp == nil || resp.StatusCode != 200 {
		e := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "Code is invalid!")
		return v_api.WriteError(ctx, e)
	}

	err := api.regionsController.UpdateRegions(&payload)
	if err != nil {
		var e *v_proto.VolioRpcError
		if strings.Contains(err.Error(), "regions_name_key") {
			e = v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Name is %s already exist!", payload.Name))
		} else if strings.Contains(err.Error(), "regions_code_key") {
			e = v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Code is %s already exist!", payload.Code))
		} else {
			e = v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error())
		}
		return v_api.WriteError(ctx, e)
	}
	return v_api.WriteSuccessEmptyContent(ctx)

}

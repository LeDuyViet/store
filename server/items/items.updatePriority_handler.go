package items_handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *ItemsAPI) UpdatePriority(ctx *fiber.Ctx) error {
	storeId := ctx.Params("store_id")
	storeIdInt, err := strconv.Atoi(storeId)
	if err != nil {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	}

	region := ctx.Query("region", "")

	payload := struct {
		Data []*do.ItemsPublicDo `json:"data"`
	}{}

	payload.Data = make([]*do.ItemsPublicDo, 0)

	if region == "" {
		resp := store.GetStoreClient().Forward(ctx)
		return ctx.Send(resp)
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error()))
	}

	priorities := make(map[int8]bool, 0)
	for _, item := range payload.Data {
		if api.itemsController.CheckItemExist(item.ID, int32(storeIdInt)) {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Item %v invalid ", item.ID)))
		}
		if ok := priorities[item.Priority]; ok == true {
			return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), fmt.Sprintf("Priorty %v is duplicate ", item.Priority)))
		} else {
			priorities[item.Priority] = true
		}
	}
	err = api.itemsController.UpdatePriority(int32(storeIdInt), payload.Data)

	if err != nil {
		// err := fmt.Errorf("modules.UpdatePriority: %v")
		return store_api.WriteError(ctx, err)
	}

	return store_api.WriteSuccess(ctx, nil)
}

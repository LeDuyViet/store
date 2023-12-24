package apps

import (
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (ctrl *AppsController) AdminGet(id string) interface{} {
	app := ctrl.dao.appsDAO.GetApp(id)
	if app != nil {
		return app
	}

	return nil
}

func (ctrl *AppsController) AdminGets(offset, limit int) interface{} {
	return ctrl.dao.appsDAO.GetAll(offset, limit)
}

func (ctrl *AppsController) AdminCreate(reqInterface interface{}) *v_proto.VolioRpcError {
	req := reqInterface.(*do.CreateAppReq)
	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to create")
	}

	timeNow := int32(time.Now().Unix())
	md := &models.AppMD{
		ID:          uuid.NewString(),
		Name:        req.Name,
		PackageName: req.PackageName,
		Photo:       req.Photo,
		StoreID:     req.StoreID,
		CreatedTime: timeNow,
		UpdatedTime: timeNow,
	}

	req.ID = md.ID

	if err := ctrl.dao.appsDAO.InsertApp(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *AppsController) AdminUpdate(reqInterface interface{}) *v_proto.VolioRpcError {
	req := reqInterface.(*do.UpdateAppReq)
	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to update")
	}

	md := ctrl.GetApp(req.ID)
	if md == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_NOT_FOUND), "app not found")
	}

	md.Name = req.Name
	md.PackageName = req.PackageName
	md.Photo = req.Photo

	timeNow := int32(time.Now().Unix())
	md.UpdatedTime = timeNow

	if err := ctrl.dao.appsDAO.Update(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}
func (ctrl *AppsController) AdminDelete(id string) *v_proto.VolioRpcError {
	if err := ctrl.dao.appsDAO.Delete(id); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), "can not delete store")
	}

	return nil
}

package apps

import (
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

var _ core.Admin = (*AppsController)(nil)
var _ core.AppsCallback = (*AppsController)(nil)

func (ctrl *AppsController) GetApp(id string) *models.AppMD {
	app := ctrl.dao.appsDAO.GetApp(id)
	if app != nil {
		return app
	}

	return nil
}

func (ctrl *AppsController) GetByStore(Admin string, offset, limit int) []*models.AppMD {
	return ctrl.dao.appsDAO.GetByStore(Admin, offset, limit)
}

func (ctrl *AppsController) GetAllByStore(Admin string) []*models.AppMD {
	return ctrl.dao.appsDAO.GetAllAppsByStoreID(Admin)
}

func (ctrl *AppsController) CreateApp(req *do.CreateAppReq) *v_proto.VolioRpcError {
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

	if err := ctrl.dao.appsDAO.InsertApp(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

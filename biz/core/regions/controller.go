package regions

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type regionDAO struct {
	regionDAO    *postgres_dao.RegionsDAO
	AppRegionDAO *postgres_dao.AppRegionsDAO
}

type RegionsController struct {
	dao *regionDAO
}

func (ctrl *RegionsController) InstallController() {
	ctrl.dao.regionDAO = dao.GetRegionsDAO(dao.STORES_DB_MASTER)
	ctrl.dao.AppRegionDAO = dao.GetAppRegionsDAO(dao.STORES_DB_MASTER)
}

func (ctrl *RegionsController) RegisterCallback(cb interface{}) {

}

func (ctrl *RegionsController) AfterInstalledDone() {
	v_log.V(3).Infof("RegionsController::AfterInstalledDone")
}

func init() {
	f_core.RegisterCoreController(&RegionsController{dao: &regionDAO{}})
}

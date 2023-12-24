package regions

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (ctrl *RegionsController) GetAllRegions() []*models.RegionMD {
	return ctrl.dao.regionDAO.GetRegions()
}

func (ctrl *RegionsController) InsertRegions(regions *models.RegionMD) error {
	timeNow := int32(time.Now().Unix())
	regions.ID = uuid.New().String()
	regions.Code = strings.ToUpper(regions.Code)
	regions.CreatedTime = timeNow
	regions.UpdatedTime = timeNow
	return ctrl.dao.regionDAO.InsertRegions(regions)
}

func (ctrl *RegionsController) UpdateRegions(regions *models.RegionMD) error {
	result := ctrl.dao.regionDAO.CheckRegionExist(regions.ID)
	if !result {
		return fmt.Errorf("Region with id is %d not exist!", regions.ID)
	}
	regions.Code = strings.ToUpper(regions.Code)
	regions.UpdatedTime = int32(time.Now().Unix())
	return ctrl.dao.regionDAO.UpdateRegion(regions)
}

func (ctrl *RegionsController) CheckRegionExist(condition string) bool {
	return ctrl.dao.regionDAO.CheckRegionExist(condition)
}

func (ctrl *RegionsController) AddApp(appRegion *models.AppRegionMD) *v_proto.VolioRpcError {
	timeNow := int32(time.Now().Unix())
	appRegionMd := &models.AppRegionMD{
		RegionID:    appRegion.RegionID,
		AppID:       appRegion.AppID,
		CreatedTime: timeNow,
		UpdatedTime: timeNow,
	}

	if ctrl.dao.AppRegionDAO.CheckAppRegionExist(appRegionMd) {
		err := fmt.Sprintf("region_id is %s and app_id is %s already exist ", appRegionMd.RegionID, appRegionMd.AppID)
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err)
	}

	err := ctrl.dao.AppRegionDAO.InsertAppRegion(appRegionMd)
	if err != nil {
		err := fmt.Sprintf("region_id is %s and app_id is %s already exist ", appRegionMd.RegionID, appRegionMd.AppID)
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err)
	}

	return nil
}

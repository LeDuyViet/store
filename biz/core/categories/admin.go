package categories

import (
	"time"

	"github.com/google/uuid"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

var _ core.Admin = (*CategoriesController)(nil)

func (ctrl *CategoriesController) AdminGet(id string) interface{} {

	if cate := ctrl.dao.categoriesDAO.GetCategory(id); cate != nil {
		do := &do.CategoriesDO{
			ID:       cate.ID,
			Name:     cate.Name,
			ParentID: cate.ParentID,
			Priority: cate.Priority,
			Icon:     cate.Icon,
			Status:   cate.Status,
			IsPro:    cate.IsPro,
		}

		do.Childrens = ctrl.GetByParent(cate.ID)

		return do
	}

	return nil
}

func (ctrl *CategoriesController) AdminGets(offset, limit int) interface{} {
	return ctrl.dao.categoriesDAO.GetCategories(offset, limit)
}

func (ctrl *CategoriesController) AdminCreate(reqInterface interface{}) *v_proto.VolioRpcError {
	req := reqInterface.(*do.CreateCategoryReq)
	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to create")
	}

	timeNow := int32(time.Now().Unix())
	md := &models.CategoryMD{
		ID:          uuid.NewString(),
		Name:        req.Name,
		ParentID:    req.ParentID,
		Priority:    req.Priority,
		Icon:        req.Icon,
		Status:      req.Status,
		IsPro:       req.IsPro,
		AppID:       req.AppID,
		CreatedTime: timeNow,
		UpdatedTime: timeNow,
	}

	if md.ParentID == "" {
		md.ParentID = md.ID
	}

	req.ID = md.ID

	if err := ctrl.dao.categoriesDAO.InsertCategory(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *CategoriesController) AdminUpdate(reqInterface interface{}) *v_proto.VolioRpcError {
	req := reqInterface.(*do.UpdateCategoryReq)
	if req == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "not found object to update")
	}

	md := ctrl.GetCategory(req.ID)
	if md == nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_NOT_FOUND), "app not found")
	}

	md.Name = req.Name
	md.Priority = req.Priority
	md.Icon = req.Icon
	md.Status = req.Status
	md.IsPro = req.IsPro

	timeNow := int32(time.Now().Unix())
	md.UpdatedTime = timeNow

	// TODO: sort all with priority

	if err := ctrl.dao.categoriesDAO.Update(md); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *CategoriesController) AdminDelete(id string) *v_proto.VolioRpcError {
	// Lấy danh sách các category con
	ids := []string{id}
	children := ctrl.GetByParent(id)

	ids = append(ids, ctrl.getAllChildrenID(children)...)

	if err := ctrl.dao.categoriesDAO.DeleteMultiple(ids); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *CategoriesController) getAllChildrenID(children []*do.CategoriesDO) []string {
	ids := []string{}
	for _, child := range children {
		ids = append(ids, child.ID)
		ids = append(ids, ctrl.getAllChildrenID(child.Childrens)...)
	}

	return ids
}

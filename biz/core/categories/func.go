package categories

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

var _ core.CategoriesCallback = (*CategoriesController)(nil)

func (ctrl *CategoriesController) CreateCategory(category *models.CategoryMD) *v_proto.VolioRpcError {
	if err := ctrl.dao.categoriesDAO.InsertCategory(category); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *CategoriesController) GetCategory(id string) *models.CategoryMD {
	return ctrl.dao.categoriesDAO.GetCategory(id)
}

// func (ctrl *CategoriesController) ListActive(module *models.DataMigratesMD, perPage, current_page int) *do.ResPaginate {

// 	data := []*do.CategoriesPublicDo{}
// 	categories, total := ctrl.dao.categoriesDAO.GetCategorybyModuleId(module.MigrationID, perPage, current_page)
// 	if total == 0 {
// 		return nil
// 	}
// 	for _, category := range categories {

// 		categoryMigrate := ctrl.dataMigratesCallback.GetByID(category.ID)
// 		if categoryMigrate == nil {
// 			continue
// 		}

// 		category.IdInt = categoryMigrate.DataID
// 		category.ModuleIDInt = module.DataID
// 		custom_fields := ctrl.customFieldCallback.GetForPublicCategory(category.ID, category.ModuleID)
// 		categoryPublic := do.CreateCategoriesPublicDo(category, custom_fields)

// 		data = append(data, categoryPublic)
// 	}

// 	return &do.ResPaginate{
// 		Data:  data,
// 		Links: do.CreateLinksDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr), fmt.Sprintf("modules=%v", module.DataID)),
// 		Meta:  do.CreateMetaDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr)),
// 	}
// }

// func (ctrl *CategoriesController) GetCategoriesByModuleIdForAdmin(module *models.DataMigratesMD, perPage, current_page int) *do.ResPaginate {

// 	data := []*do.CategoriesPublicDo{}
// 	Categories, total := ctrl.dao.categoriesDAO.GetCategorybyModuleIdForAdmin(module.MigrationID, perPage, current_page)
// 	if total == 0 {
// 		return nil
// 	}
// 	for _, category := range Categories {

// 		categoryMigrate := ctrl.dataMigratesCallback.GetByID(category.ID)
// 		if categoryMigrate == nil {
// 			continue
// 		}

// 		category.IdInt = categoryMigrate.DataID
// 		category.ModuleIDInt = module.DataID
// 		custom_fields := ctrl.customFieldCallback.GetForPublicCategory(category.ID, category.ModuleID)
// 		categoryPublic := do.CreateCategoriesPublicDo(category, custom_fields)
// 		data = append(data, categoryPublic)
// 	}

// 	return &do.ResPaginate{
// 		Data:  data,
// 		Links: do.CreateLinksDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr), fmt.Sprintf("modules=%v", module.DataID)),
// 		Meta:  do.CreateMetaDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr)),
// 	}
// }

func (ctrl *CategoriesController) UpdatePriority(categories []*models.CategoryMD) error {
	var err error
	for _, category := range categories {
		err = ctrl.dao.categoriesDAO.UpdatePriority(category.ID, category.Priority)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctrl *CategoriesController) CheckCategoryExist(categoryId string) bool {
	return ctrl.dao.categoriesDAO.CheckCategoryExist(categoryId)
}

func (ctrl *CategoriesController) GetByApp(appID string, offset, limit int, onlyRoot ...bool) []*do.CategoriesDO {
	rootCates := ctrl.dao.categoriesDAO.GetByApp(appID, offset, limit, true)

	if len(onlyRoot) > 0 && onlyRoot[0] {
		return rootCates
	}

	for i := 0; i < len(rootCates); i++ {
		childs := ctrl.GetByParent(rootCates[i].ID)
		rootCates[i].Childrens = childs
	}

	return rootCates
}

func (ctrl *CategoriesController) GetByParent(parentID string) []*do.CategoriesDO {
	categories := ctrl.dao.categoriesDAO.GetByParent(parentID)
	for i := 0; i < len(categories); i++ {
		childs := ctrl.GetByParent(categories[i].ID)
		categories[i].Childrens = childs
	}

	// for _, cate := range categories {
	// 	ctf :=  ctrl.customFieldsCallback.GetByTableable(cate.ID, models.KTypeCategory)
	// 	cate.CustomFields =
	// }

	return categories
}

func (ctrl *CategoriesController) GetAllByApp(appID string) []*models.CategoryMD {
	return ctrl.dao.categoriesDAO.GetAllByApp(appID)
}

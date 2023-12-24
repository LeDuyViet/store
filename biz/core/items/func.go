package items

import (
	"fmt"

	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/env"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

var _ core.ItemsCallback = (*ItemsController)(nil)

func (ctrl *ItemsController) GetItem(id string) *models.ItemMD {
	return ctrl.dao.itemsDAO.GetItem(id)
}

func (ctrl *ItemsController) CreateItem(item *models.ItemMD) *v_proto.VolioRpcError {
	if err := ctrl.dao.itemsDAO.InsertItem(item); err != nil {
		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
	}

	return nil
}

func (ctrl *ItemsController) ListActive(category *models.DataMigratesMD, perPage, current_page int) *do.ResPaginate {
	data := []*do.ItemsPublicDo{}
	items, total := ctrl.dao.itemsDAO.GetItemsbyCategoryId(category.MigrationID, perPage, current_page)
	if total == 0 {
		return nil
	}

	for _, item := range items {
		itemMigrate := ctrl.dataMigratesCallback.GetByID(item.ID)
		if itemMigrate == nil {
			continue
		}

		item.IdInt = itemMigrate.DataID
		item.CategoryIDInt = category.DataID
		custom_fields := ctrl.customFieldCallback.GetForPublicItem(item.ID, item.CategoryID)
		itemPublic := do.CreateItemsPublicDo(item, custom_fields)

		data = append(data, itemPublic)
	}

	return &do.ResPaginate{
		Data:  data,
		Links: do.CreateLinksDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr), fmt.Sprintf("categories=%v", category.DataID)),
		Meta:  do.CreateMetaDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr)),
	}
}

func (ctrl *ItemsController) UpdatePriority(storeId int32, items []*do.ItemsPublicDo) error {
	var err error
	for _, item := range items {

		itemMigrate := ctrl.dataMigratesCallback.GetByDataMigrate(&models.DataMigratesMD{
			StoreID:  storeId,
			DataID:   item.ID,
			DataType: models.KTypeCategory,
		})

		err = ctrl.dao.itemsDAO.UpdatePriority(itemMigrate.MigrationID, int(item.Priority))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctrl *ItemsController) CheckItemExist(itemId, storeId int32) bool {
	itemMigrate := ctrl.dataMigratesCallback.GetByDataMigrate(&models.DataMigratesMD{
		StoreID:  storeId,
		DataID:   itemId,
		DataType: models.KTypeCategory,
	})
	return ctrl.dao.itemsDAO.CheckItemExist(itemMigrate.MigrationID)
}

func (ctrl *ItemsController) GetItemsbyCategoryIdForAdmin(category *models.DataMigratesMD, perPage, current_page int) *do.ResPaginate {
	data := []*do.ItemsPublicDo{}
	items, total := ctrl.dao.itemsDAO.GetItemsbyCategoryIdForAdmin(category.MigrationID, perPage, current_page)
	if total == 0 {
		return nil
	}

	for _, item := range items {

		itemMigrate := ctrl.dataMigratesCallback.GetByID(item.ID)
		if itemMigrate == nil {
			continue
		}

		item.IdInt = itemMigrate.DataID
		item.CategoryIDInt = category.DataID
		custom_fields := ctrl.customFieldCallback.GetForPublicItem(item.ID, item.CategoryID)
		itemPublic := do.CreateItemsPublicDo(item, custom_fields)
		data = append(data, itemPublic)
	}

	return &do.ResPaginate{
		Data:  data,
		Links: do.CreateLinksDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr), fmt.Sprintf("categories=%v", category.DataID)),
		Meta:  do.CreateMetaDo(current_page, perPage, total, fmt.Sprintf(basePath, env.Addr)),
	}
}

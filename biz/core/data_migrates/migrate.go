package data_migrates

// import (
// 	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
// 	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
// 	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
// 	"gitlab.volio.vn/tech/fountain/baselib/base"
// 	"gitlab.volio.vn/tech/fountain/baselib/v_log"
// )

// func (ctrl *DataMigratesController) migrateStore(client *store.Client, storeId int32) *models.StoreMD {
// 	storeRes := client.GetStoresPrivate()

// 	for _, store := range storeRes {
// 		if store.ID == storeId {
// 			storeM, ok := ctrl.CheckExistInMigration(models.KTypeStoreData, store.ID, storeId)
// 			if ok {
// 				store := ctrl.dao.storesDAO.GetStore(storeM.MigrationID)
// 				if store != nil {
// 					return store
// 				}
// 				// if exist in data migration but not in new store => insert it again
// 			}

// 			storeMd := store.ToModel()
// 			storeMd.ID = storeM.MigrationID
// 			storeMd.CreatedTime = storeM.CreatedTime
// 			storeMd.UpdatedTime = storeM.UpdatedTime
// 			err := ctrl.dao.storesDAO.InsertStore(storeMd)
// 			if err != nil {
// 				v_log.V(3).Infof("InsertStore - Error: %+v", err)
// 			}

// 			return storeMd
// 		}
// 	}

// 	return nil
// }

// func (ctrl *DataMigratesController) migrateApps(client *store.Client, appId, storeId int32, storeIdMigration string) *models.AppMD {
// 	appRes := client.GetAppsPrivate(storeId)

// 	for _, app := range appRes {
// 		if app.ID == appId {
// 			app.StoreID = storeId
// 			appMigration, ok := ctrl.CheckExistInMigration(models.KTypeAppData, app.ID, storeId)
// 			if ok {
// 				app := ctrl.dao.appsDAO.GetApp(appMigration.MigrationID)
// 				if app != nil {
// 					return app
// 				}
// 			}

// 			appMd := app.ToModel()
// 			appMd.ID = appMigration.MigrationID
// 			appMd.StoreID = storeIdMigration
// 			appMd.CreatedTime = appMigration.CreatedTime
// 			appMd.UpdatedTime = appMigration.UpdatedTime
// 			err := ctrl.dao.appsDAO.InsertApp(appMd)
// 			if err != nil {
// 				v_log.V(3).Infof("InsertApp - Error: %+v", err)
// 			}

// 			return appMd
// 		}
// 	}

// 	return nil
// }

// func (ctrl *DataMigratesController) migrateCustomFields(client *store.Client, appId, storeId int32, appIdMigration string) []*models.CustomFieldMD {
// 	customFieldRes := client.GetCustomFieldsPrivate(storeId, appId)

// 	var customFields []*models.CustomFieldMD

// 	for _, customField := range customFieldRes {
// 		customField.AppID = appId
// 		customFieldMigration, ok := ctrl.CheckExistInMigration(models.KTypeCustomField, customField.ID, storeId)
// 		if ok {
// 			customField := ctrl.dao.customFieldsDAO.GetCustomField(customFieldMigration.MigrationID)
// 			if customField != nil {
// 				customFields = append(customFields, customField)
// 				continue
// 			}
// 		}

// 		customFieldMd := customField.ToModel()
// 		customFieldMd.ID = customFieldMigration.MigrationID
// 		customFieldMd.AppID = appIdMigration
// 		customFieldMd.CreatedTime = customFieldMigration.CreatedTime
// 		customFieldMd.UpdatedTime = customFieldMigration.UpdatedTime
// 		err := ctrl.dao.customFieldsDAO.InsertCustomField(customFieldMd)
// 		if err != nil {
// 			v_log.V(3).Infof("InsertCustomField - Error: %+v", err)
// 		}

// 		customFields = append(customFields, customFieldMd)
// 	}

// 	return customFields
// }

// func (ctrl *DataMigratesController) migrateCategories(client *store.Client, moduleId, storeId int32, moduleIdMigration string) []*models.CategoryMD {
// 	categoryRes := client.GetCategoriesPrivate(storeId, moduleId)

// 	var categories []*models.CategoryMD
// 	for _, categoryDo := range categoryRes {
// 		categoryDo.ModuleID = moduleId
// 		categoryMigration, ok := ctrl.CheckExistInMigration(models.KTypeCategory, categoryDo.ID, storeId)

// 		ctrl.migrateCustomFieldTables(categoryDo.CustomFieldTables, models.KTypeCategory, storeId, categoryMigration)

// 		if ok {
// 			category := ctrl.dao.categoriesDAO.GetCategory(categoryMigration.MigrationID)
// 			if category != nil {
// 				category.DataMigratesMD = categoryMigration
// 				categories = append(categories, category)
// 				continue
// 			}
// 		}

// 		categoryMd := categoryDo.ToModel()
// 		categoryMd.ID = categoryMigration.MigrationID
// 		categoryMd.ModuleID = moduleIdMigration
// 		categoryMd.CreatedTime = categoryMigration.CreatedTime
// 		categoryMd.UpdatedTime = categoryMigration.UpdatedTime

// 		categoryMd.DataMigratesMD = categoryMigration

// 		err := ctrl.dao.categoriesDAO.InsertCategory(categoryMd)
// 		if err != nil {
// 			v_log.V(3).Errorf("InsertCategory - Error: %+v", err)
// 		}

// 		categories = append(categories, categoryMd)
// 	}

// 	return categories
// }

// func (ctrl *DataMigratesController) migrateCustomFieldTables(CustomFieldTables []*do.StoreCustomFieldTableDO, dataType string, storeId int32, moduleMigration *models.DataMigratesMD) []*models.CustomFieldDataMD {
// 	var customFieldTables []*models.CustomFieldDataMD

// 	for _, customFieldTable := range CustomFieldTables {
// 		customFieldTableMigration, ok := ctrl.CheckExistInMigration(models.KTypeCustomFieldTable, customFieldTable.ID, storeId)
// 		if !ok {
// 			customFieldMigration := ctrl.GetByDataMigrate(&models.DataMigratesMD{
// 				StoreID:  storeId,
// 				DataID:   customFieldTable.CustomFieldID,
// 				DataType: models.KTypeCustomField,
// 			})

// 			customFieldTablesMd := &models.CustomFieldDataMD{
// 				ID:                       customFieldTableMigration.MigrationID,
// 				CustomFieldID:            customFieldMigration.MigrationID,
// 				Value:                    customFieldTable.CustomFieldValue,
// 				CustomFieldTableableID:   moduleMigration.MigrationID,
// 				CustomFieldTableableType: dataType,
// 				IsActive:                 base.Int8ToBool(customFieldTable.IsActive),
// 				CreatedTime:              0,
// 				UpdatedTime:              0,
// 			}
// 			if err := ctrl.dao.customFieldTablesDAO.Insert(customFieldTablesMd); err != nil {
// 				v_log.V(3).Errorf("DataMigratesController::migrateCustomFieldTables - Error when insert custom field table: %+v", err)
// 			}

// 			customFieldTables = append(customFieldTables, customFieldTablesMd)
// 		}

// 		customFieldTable := ctrl.dao.customFieldTablesDAO.Get(customFieldTableMigration.MigrationID)
// 		customFieldTables = append(customFieldTables, customFieldTable)
// 	}

// 	return customFieldTables
// }

// func (ctrl *DataMigratesController) migrateItems(client *store.Client, categoryIdRaw, storeIdRaw int32, categoryIdMigration string) []*models.ItemMD {
// 	itemRes := client.GetItemsPrivate(storeIdRaw, categoryIdRaw)

// 	var items []*models.ItemMD
// 	for _, itemDo := range itemRes {
// 		itemDo.CategoryID = categoryIdRaw
// 		itemMigration, ok := ctrl.CheckExistInMigration(models.KTypeItem, itemDo.ID, storeIdRaw)

// 		ctrl.migrateCustomFieldTables(itemDo.CustomFieldTables, models.KTypeItem, storeIdRaw, itemMigration)

// 		if ok {
// 			item := ctrl.dao.itemsDAO.GetItem(itemMigration.MigrationID)
// 			if item != nil {
// 				items = append(items, item)
// 				continue
// 			}
// 		}

// 		itemMd := itemDo.ToModel()
// 		itemMd.ID = itemMigration.MigrationID
// 		itemMd.CategoryID = categoryIdMigration
// 		itemMd.CreatedTime = itemMigration.CreatedTime
// 		itemMd.UpdatedTime = itemMigration.UpdatedTime

// 		err := ctrl.dao.itemsDAO.InsertItem(itemMd)
// 		if err != nil {
// 			v_log.V(3).Errorf("InsertItem - Error: %+v", err)
// 		}

// 		items = append(items, itemMd)
// 	}

// 	return items
// }

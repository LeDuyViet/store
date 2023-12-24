package data_migrates

// import (
// 	"fmt"
// 	"time"
// 	"unsafe"

// 	"github.com/google/uuid"
// 	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
// 	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
// 	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
// 	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/do"
// 	"gitlab.volio.vn/tech/fountain/baselib/v_log"
// 	"gitlab.volio.vn/tech/fountain/proto/v_proto"
// )

// var _ core.DataMigratesCallback = (*DataMigratesController)(nil)

// func (ctrl *DataMigratesController) GetByDataMigrate(dataMigrate *models.DataMigratesMD) *models.DataMigratesMD {
// 	v, ok := ctrl.lruCache.Get(models.CreateDataMigrageKeyWithParam(dataMigrate))

// 	if ok {
// 		if model, o := v.(*LruCache); o && model != nil {
// 			return model.DataMigratesMD
// 		}
// 	}

// 	model := ctrl.dao.dataMigratesDAO.GetByMigrateData(dataMigrate)
// 	if model != nil {
// 		ctrl.lruCache.Set(models.CreateDataMigrageKeyWithParam(dataMigrate), &LruCache{model, int(unsafe.Sizeof(model))})
// 	}

// 	// v2, ok2 := ctrl.lruCache.Get(models.CreateDataMigrageKeyWithParam(dataMigrate))

// 	// if ok2 {
// 	// 	if model, o := v2.(*LruCache); o && model != nil {
// 	// 		if model != nil && model.DataMigratesMD == nil {
// 	// 			allcache := ctrl.lruCache.Items()
// 	// 			log.Printf("allcache: %+v", allcache)
// 	// 		}
// 	// 	}
// 	// }

// 	return model
// }

// func (ctrl *DataMigratesController) CreateDataMigrate(dataMigrate *models.DataMigratesMD) *v_proto.VolioRpcError {
// 	if err := ctrl.dao.dataMigratesDAO.InsertDataMigrate(dataMigrate); err != nil {
// 		return v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_SAVE_DATA), err.Error())
// 	}

// 	return nil
// }

// func (ctrl *DataMigratesController) GetByID(id string) *models.DataMigratesMD {
// 	v, ok := ctrl.lruCache.Get(models.CreateDataMigrageKey(id))

// 	if ok {
// 		if model, o := v.(*LruCache); o && model != nil {
// 			return model.DataMigratesMD
// 		}
// 	}

// 	model := ctrl.dao.dataMigratesDAO.GetByID(id)
// 	if model != nil {
// 		ctrl.lruCache.Set(models.CreateDataMigrageKey(id), &LruCache{model, int(unsafe.Sizeof(model))})
// 	}

// 	return model
// }

// // if data not exist in migration, insert it and return
// func (ctrl *DataMigratesController) CheckExistInMigration(dataType string, dataId, storeId int32) (*models.DataMigratesMD, bool) {
// 	timeNow := int32(time.Now().Unix())

// 	storeMigrate := &models.DataMigratesMD{
// 		StoreID:  storeId,
// 		DataID:   dataId,
// 		DataType: dataType,
// 	}

// 	dataMigrate := ctrl.GetByDataMigrate(storeMigrate)
// 	if dataMigrate != nil {
// 		return dataMigrate, true
// 	}

// 	storeMigrate.MigrationID = uuid.New().String()
// 	storeMigrate.StoreID = storeId
// 	storeMigrate.IsSync = true
// 	storeMigrate.CreatedTime = timeNow
// 	storeMigrate.UpdatedTime = timeNow

// 	err := ctrl.CreateDataMigrate(storeMigrate)
// 	if err != nil {
// 		v_log.V(3).Infof("InsertDataMigrate - Error: %+v", err)
// 	}

// 	return storeMigrate, false
// }

// func (ctrl *DataMigratesController) MigrateData(storeId, appId int32) {
// 	ctrl.locker.Lock()
// 	client := store.GetStoreClient()
// 	client.Token = client.Login().AccessToken

// 	store := ctrl.migrateStore(client, storeId)
// 	if store == nil {
// 		err := fmt.Errorf("DataMigratesController::MigrateData - StoreId: %d is not exist", storeId)
// 		v_log.V(3).Errorf("DataMigratesController::MigrateData - Error: %+v", err)
// 	}

// 	if store == nil {
// 		v_log.V(3).Errorf("DataMigratesController::MigrateData - StoreId: %d is not exist", storeId)
// 		return
// 	}

// 	app := ctrl.migrateApps(client, appId, storeId, store.ID)

// 	customFields := ctrl.migrateCustomFields(client, appId, storeId, app.ID)
// 	v_log.V(3).Infof("DataMigratesController::MigrateData - CustomFields: %+v", customFields)

// 	categories := ctrl.migrateCategories(client, appId, storeId, app.ID)
// 	for _, category := range categories {
// 		ctrl.migrateItems(client, category.DataMigratesMD.DataID, storeId, category.ID)
// 	}

// 	ctrl.locker.Unlock()
// }

// func (ctrl *DataMigratesController) MigrateStores() {
// 	ctrl.locker.Lock()
// 	client := store.GetStoreClient()
// 	client.Token = client.Login().AccessToken

// 	storesMigrate := ctrl.dao.storesDAO.GetAll()
// 	stores := client.GetStoresPrivate()

// 	storeNotMigrates := make([]*do.StoreDO, 0)

// 	mStore := make(map[string]bool, 0)
// 	for _, m := range storesMigrate {
// 		mStore[m.Address] = true
// 	}

// 	for _, store := range stores {
// 		if ok := mStore[store.Address]; !ok {
// 			storeNotMigrates = append(storeNotMigrates, store)
// 		}
// 	}

// 	for _, storeNotMigrate := range storeNotMigrates {
// 		ctrl.migrateStore(client, storeNotMigrate.ID)
// 	}

// 	ctrl.locker.Unlock()
// }

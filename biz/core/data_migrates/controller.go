package data_migrates

import (
	"sync"

	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/baselib/cache/lru_cache"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type dataMigratesDAO struct {
	dataMigratesDAO      *postgres_dao.DataMigratesDAO
	storesDAO            *postgres_dao.StoresDAO
	appsDAO              *postgres_dao.AppsDAO
	customFieldsDAO      *postgres_dao.CustomFieldsDAO
	customFieldTablesDAO *postgres_dao.CustomFieldDataDAO
	categoriesDAO        *postgres_dao.CategoriesDAO
	itemsDAO             *postgres_dao.ItemsDAO
}

type DataMigratesController struct {
	dao *dataMigratesDAO

	lruCache *lru_cache.LRUCache

	locker sync.Mutex
}

type LruCache struct {
	*models.DataMigratesMD
	size int
}

func (lc *LruCache) Size() int {
	return lc.size
}

func (ctrl *DataMigratesController) InstallController() {
	ctrl.lruCache = lru_cache.NewLRUCache(10 * 1024 * 1024)

	ctrl.dao.dataMigratesDAO = dao.GetDataMigratesDAO(dao.STORES_DB_MASTER)
	ctrl.dao.storesDAO = dao.GetStoresDAO(dao.STORES_DB_MASTER)
	ctrl.dao.appsDAO = dao.GetAppsDAO(dao.STORES_DB_MASTER)
	ctrl.dao.categoriesDAO = dao.GetCategoriesDAO(dao.STORES_DB_MASTER)
	ctrl.dao.itemsDAO = dao.GetItemsDAO(dao.STORES_DB_MASTER)
	ctrl.dao.customFieldsDAO = dao.GetCustomFieldsDAO(dao.STORES_DB_MASTER)
}

func (ctrl *DataMigratesController) RegisterCallback(cb interface{}) {
}

func (ctrl *DataMigratesController) AfterInstalledDone() {
	v_log.V(3).Infof("DataMigratesController::AfterInstalledDone")
}

func init() {
	f_core.RegisterCoreController(&DataMigratesController{dao: &dataMigratesDAO{}})
}

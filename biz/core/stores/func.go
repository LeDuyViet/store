package stores

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
)

var _ core.Admin = (*StoresController)(nil)

var _ core.StoresCallback = (*StoresController)(nil)

func (ctrl *StoresController) GetStore(id string) *models.StoreMD {
	return ctrl.dao.storesDAO.GetStore(id)
}

func (ctrl *StoresController) GetAll() []*models.StoreMD {
	return ctrl.dao.storesDAO.GetAll()
}

func (ctrl *StoresController) AdminGets(offset, limit int) interface{} {
	return ctrl.dao.storesDAO.GetStores(offset, limit)
}

func (ctrl *StoresController) GetByPackageName(packageName string) []*models.StoreMD {
	return ctrl.dao.storesDAO.GetByPackageName(packageName)
}

func (ctrl *StoresController) GetStoresByAddress(address string) *models.StoreMD {
	v := ctrl.cache.Get(address)

	if v != nil {
		if store, o := v.(*models.StoreMD); o && store != nil {
			return store
		}
	}

	store := ctrl.dao.storesDAO.GetStoresByAddress(address)
	if store != nil {
		ctrl.cache.Put(address, store, 0)
	}

	return store
}

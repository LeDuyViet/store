/* !!
 * File: controller.go
 * File Created: Tuesday, 3rd January 2023 3:36:33 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Tuesday, 3rd January 2023 3:38:06 pm
 * Modified By: KimEricko™ (phamkim.pr@gmail.com>)
 * -----
 * Copyright 2022 - 2023 Volio, Volio Vietnam
 * All rights reserved.
 *
 * Licensed under the GNU GENERAL PUBLIC LICENSE, Version 3.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  https://www.gnu.org/licenses/gpl-3.0.html
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Developer: NhokCrazy199 (phamkim.pr@gmail.com)
 * -----
 * HISTORY:
 * Date      	By	Comments
 * ----------	---	---------------------------------------------------------
 */

package stores

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/admin"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/fountain/baselib/cache"
	_ "gitlab.volio.vn/tech/fountain/baselib/cache/memory_cache"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	f_core "gitlab.volio.vn/tech/fountain/biz/core"
)

type storesDAO struct {
	storesDAO *postgres_dao.StoresDAO
}

type StoresController struct {
	dao   *storesDAO
	cache cache.Cache

	*admin.AdminController

	appsCallback         core.AppsCallback
	dataMigratesCallback core.DataMigratesCallback
}

func (ctrl *StoresController) InstallController() {
	memoryCache, err := cache.NewCache("memory", "")
	if err != nil {
		panic(err)
	}
	ctrl.cache = memoryCache

	ctrl.dao.storesDAO = dao.GetStoresDAO(dao.STORES_DB_MASTER)
}

func (ctrl *StoresController) RegisterCallback(cb interface{}) {
	switch x := cb.(type) {
	case core.DataMigratesCallback:
		ctrl.dataMigratesCallback = x
	case core.AppsCallback:
		ctrl.appsCallback = x
	case *admin.AdminController:
		ctrl.AdminController = x

	}
}

func (ctrl *StoresController) AfterInstalledDone() {
	v_log.V(3).Infof("StoresController::AfterInstalledDone")
	ctrl.registerCustomValidorRule()
}

func init() {
	f_core.RegisterCoreController(&StoresController{dao: &storesDAO{}})
}

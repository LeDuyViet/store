/* !!
 * File: server.go
 * File Created: Thursday, 3rd November 2022 4:09:35 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd November 2022 4:09:36 pm
 * Modified By: KimEricko™ (phamkim.pr@gmail.com>)
 * -----
 * Copyright 2022 - 2022 Volio, Volio Vietnam
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

package server

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/stores"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	apps_handler "gitlab.volio.vn/tech/backend/store-pro/server/apps"
	categories_handler "gitlab.volio.vn/tech/backend/store-pro/server/categories"
	items_handler "gitlab.volio.vn/tech/backend/store-pro/server/items"
	regions_handler "gitlab.volio.vn/tech/backend/store-pro/server/regions"
	stores_handler "gitlab.volio.vn/tech/backend/store-pro/server/stores"
	"gitlab.volio.vn/tech/fountain/baselib/env"
	"gitlab.volio.vn/tech/fountain/baselib/redis_client"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
	"gitlab.volio.vn/tech/fountain/biz/core/auth"
	"gitlab.volio.vn/tech/fountain/biz/core/health_check"
	f_dao "gitlab.volio.vn/tech/fountain/biz/dal/dao"

	_ "gitlab.volio.vn/tech/backend/store-pro/biz"
)

type storeProServer struct {
	apiServer *v_api.VolioAPI

	controllers []core.CoreController

	// Fountain Controller
	authController        *auth.AuthController
	healthCheckController *health_check.HealthCheckController

	storeController *stores.StoresController

	// Handler
	storesHandler     *stores_handler.StoresAPI
	appsHandler       *apps_handler.AppsAPI
	categoriesHandler *categories_handler.CategoriesAPI
	itemsHandler      *items_handler.ItemsAPI
	regionsHandler    *regions_handler.RegionAPI
}

func NewStoreProServer() *storeProServer {
	return &storeProServer{}
}

func (s *storeProServer) GetIdentification() (addr, dcName, serviceName, serverID string) {
	return env.Addr, env.DCName, env.ServiceName, env.PodName
}

func (s *storeProServer) Initialize() error {
	s.apiServer = v_api.NewVolioAPI()

	redis_client.InstallRedisClientManager()
	sql_client.InstallSQLClientManager("postgresql")

	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())
	f_dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())
	dao.InstallMysqlDAOManager(sql_client.GetSQLClientManager())

	// Install clients
	store.InstallStoreClient()

	// Install controllers
	s.controllers = core.InstallCoreControllers()

	// Install handler
	s.storesHandler = stores_handler.NewStoresAPI(s.controllers)
	s.appsHandler = apps_handler.NewAppsAPI(s.controllers)
	s.categoriesHandler = categories_handler.NewCategoriesAPI(s.controllers)
	s.itemsHandler = items_handler.NewItemsAPI(s.controllers)
	s.regionsHandler = regions_handler.NewRegionAPI(s.controllers)

	return nil

}

func (s *storeProServer) RunLoop() {
	for _, ctr := range s.controllers {
		switch controller := ctr.(type) {
		case *auth.AuthController:
			v_log.V(3).Infof("storeProServer::RunLoop installed: %T", controller)
			s.authController = controller
		case *health_check.HealthCheckController:
			v_log.V(3).Infof("storeProServer::RunLoop installed: %T", controller)
			s.healthCheckController = controller
		case *stores.StoresController:
			v_log.V(3).Infof("storeProServer::RunLoop installed: %T", controller)
			s.storeController = controller
		}
	}

	s.initRouterCurrent()
	s.initRouterV3()

	go s.apiServer.Serve()
}

func (s *storeProServer) Destroy() {
	s.apiServer.Stop()
}

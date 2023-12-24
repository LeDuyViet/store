/* !!
 * File: route.go
 * File Created: Thursday, 3rd November 2022 4:09:44 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd November 2022 4:09:44 pm
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
	"fmt"

	"gitlab.volio.vn/tech/fountain/baselib/env"
)

// Using UUID
func (s *storeProServer) initRouterV3() {
	prefix := fmt.Sprintf("/%s/api/v3.0", env.EndpointPrefix)
	apiGr := s.apiServer.Group(prefix)
	apiPrivateGr := s.apiServer.Group("/private" + prefix)
	if env.Environment != "local" {
		apiGr.Use(s.authController.JwtGateway)
	}

	_ = apiGr
	_ = apiPrivateGr

	if env.Environment == "local" {
		apiGr.Use(localMiddleGate)
		apiPrivateGr.Use(localMiddleGate)
	}

	// Stores
	{
		storeGr := apiGr.Group("/stores")
		storeGr.Get("", s.storesHandler.Gets)
		storeGr.Get("/all", s.storesHandler.GetAll)
		storeGr.Post("", s.storesHandler.Create)
		storeGr.Put("", s.storesHandler.Update)
		storeGr.Delete("/:id", s.storesHandler.Delete)

		// storeGr.Post("/regional", s.regionsHandler.InsertRegions)
		// storeGr.Put("/regional", s.storesHandler.Update)
		// storeGr.Delete("/regional", s.regionsHandler.InsertRegions)

		storeGr.Get("/:id/apps", s.storesHandler.GetApps)
		storeGr.Get("/:id/apps/all", s.storesHandler.GetAllApps)
	}

	{
		appGr := apiGr.Group("/apps")
		appGr.Get("", s.appsHandler.Gets)
		appGr.Post("", s.appsHandler.Create)
		appGr.Put("", s.appsHandler.Update)
		// appGr.Put("/order", s.appsHandler.Order)
		appGr.Delete("/:id", s.appsHandler.Delete)

		// storeGr.Post("/regional", s.regionsHandler.InsertRegions)
		// storeGr.Put("/regional", s.storesHandler.Update)
		// storeGr.Delete("/regional", s.regionsHandler.InsertRegions)

		appGr.Get("/:id/categories", s.appsHandler.GetCategories)
		appGr.Get("/:id/categories/all", s.appsHandler.GetAllCategories)
	}

	{
		categoriesGr := apiGr.Group("/categories")
		categoriesGr.Get("", s.categoriesHandler.Gets)
		categoriesGr.Post("", s.categoriesHandler.Create)
		categoriesGr.Put("", s.categoriesHandler.Update)
		categoriesGr.Delete("/:id", s.categoriesHandler.Delete)

		categoriesGr.Get("", s.categoriesHandler.Create)
		categoriesGr.Post("", s.categoriesHandler.Create)
		// storeGr.Post("/regional", s.regionsHandler.InsertRegions)
		// storeGr.Put("/regional", s.storesHandler.Update)
		// storeGr.Delete("/regional", s.regionsHandler.InsertRegions)

		categoriesGr.Get("/:id/items", s.storesHandler.Gets)
		categoriesGr.Get("/:id/items/all", s.storesHandler.Gets)
	}

	{
		itemsGr := apiGr.Group("/items")
		itemsGr.Get("", s.appsHandler.Gets)
		itemsGr.Post("", s.appsHandler.Create)
		itemsGr.Post("/multiple", s.itemsHandler.CreateMulti)
		itemsGr.Put("/:id", s.appsHandler.Update)
		// appGr.Put("/order", s.appsHandler.Order)
		itemsGr.Delete("/:id", s.appsHandler.Delete)

		// storeGr.Post("/regional", s.regionsHandler.InsertRegions)
		// storeGr.Put("/regional", s.storesHandler.Update)
		// storeGr.Delete("/regional", s.regionsHandler.InsertRegions)
	}
}

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
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store"
	"gitlab.volio.vn/tech/fountain/baselib/env"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
)

// Using id increment
func (s *storeProServer) initRouterCurrent() {
	// route dạng /:store_name/api/ thì sao (/myzgos/api/v2/public/items)
	storeProGr := s.apiServer.Group("/api/")
	storeProGrPrivateGr := s.apiServer.Group("/private/")
	if env.Environment != "local" {
		storeProGrPrivateGr.Use(s.authController.JwtGateway)
	}

	_ = storeProGr
	_ = storeProGrPrivateGr

	if env.Environment == "local" {
		storeProGr.Use(localMiddleGate)
		storeProGrPrivateGr.Use(localMiddleGate)
	}

	{
		publicGr := storeProGr.Group("/v2/public")
		publicGr.Use(s.storeController.CheckStoreIDExistence)

		// publicGr.Get("/modules", s.modulesHandler.ListActive)
		publicGr.Get("/categories", s.categoriesHandler.ListActive)
		publicGr.Get("/items", s.itemsHandler.ListActive)

		publicGr.All("*", func(c *fiber.Ctx) error {
			resp := store.GetStoreClient().Forward(c, true)
			return c.Send((resp))
		})
	}

	{
		storeGr := storeProGr.Group("/stores")
		storeGr.Get("", s.storesHandler.Gets)
		storeGr.Get("/get-all", s.storesHandler.GetAll)
		storeGr.Post("", s.storesHandler.Create)
		storeGr.Put("/:id", s.storesHandler.Update)
		storeGr.Delete("/:id", s.storesHandler.Delete)

		{
			appGr := storeGr.Group("/:store_id/apps")

			// --- New ---
			appGr.Get("", s.appsHandler.Gets)
			appGr.Get("/get-all", s.appsHandler.GetAll)
			appGr.Post("", s.appsHandler.Create)
			appGr.Put("/:id", s.appsHandler.Update)
			appGr.Delete("/:id", s.appsHandler.Delete)
			// --- --- ---

			{
				moduleGr := appGr.Group("/:app_id/modules")
				// moduleGr.Get("", s.modulesHandler.GetAll)

				// // --- New ---
				// moduleGr.Get("/:id", s.modulesHandler.Get)
				// moduleGr.Get("/search", s.modulesHandler.Search)
				// // --- --- ---

				// moduleGr.Post("", s.modulesHandler.Create)
				// moduleGr.Put("/:id", s.modulesHandler.Update)
				// moduleGr.Put("/update-priority", s.modulesHandler.UpdatePriority)

				// // --- New ---
				// moduleGr.Put("/:id", s.modulesHandler.Update)
				// moduleGr.Delete("/", s.modulesHandler.Delete)
				// --- --- ---

				{
					categoryGr := moduleGr.Group("/:module_id/categories")
					categoryGr.Get("", s.categoriesHandler.GetAll)

					// --- New ---
					categoryGr.Get("/:id", s.categoriesHandler.Gets)
					categoryGr.Post("", s.categoriesHandler.Create)
					categoryGr.Put("/:id", s.categoriesHandler.Update)
					categoryGr.Put("/update-priority", s.categoriesHandler.UpdatePriority)
					categoryGr.Delete("/:id", s.categoriesHandler.Delete)
					// --- --- ---

					{
						itemGr := categoryGr.Group("/:category_id/items")
						itemGr.Get("", s.itemsHandler.GetAll)

						// --- New ---
						itemGr.Get("/:id", s.itemsHandler.Get)
						itemGr.Post("", s.itemsHandler.Create)
						itemGr.Post("/multiple", s.itemsHandler.CreateMulti)
						itemGr.Put("/:id", s.itemsHandler.Update)
						itemGr.Put("/update-priority", s.itemsHandler.UpdatePriority)
						itemGr.Delete("/:id", s.itemsHandler.Delete)
						// --- --- ---
					}
				}
			}
		}

		{
			// --- New ---
			// customFieldGr := storeGr.Group("/:id/custom-fields")
			// customFieldGr.Get("", s.customFieldsHandler.GetAll)
			// customFieldGr.Post("", s.customFieldsHandler.Create)
			// customFieldGr.Put("/update-priority", s.customFieldsHandler.UpdatePriority)
			// customFieldGr.Delete("/:id", s.customFieldsHandler.Delete)
			// --- --- ---

		}
		storeGr.All("*", func(c *fiber.Ctx) error {
			resp := store.GetStoreClient().Forward(c)
			return c.Send((resp))
		})
	}

	{
		regionGr := storeProGr.Group("/regions")
		regionGr.Get("", s.regionsHandler.GetAllRegions)
		regionGr.Post("/create", s.regionsHandler.InsertRegions)
		regionGr.Patch("/update", s.regionsHandler.UpdateRegions)
		regionGr.Post("/add-app", s.regionsHandler.AddApp)

		// storeProGr.Get("*", s.regionsHandler.Get)
	}

}

func localMiddleGate(c *fiber.Ctx) (err error) {
	c.Locals(v_api.KContextKeyUserID, "00000000-0000-0000-0000-000000000000")
	c.Locals(v_api.KContextKeyUserName, "admin")

	return c.Next()
}

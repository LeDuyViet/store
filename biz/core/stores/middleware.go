/* !!
 * File: middleware.go
 * File Created: Wednesday, 26th July 2023 10:02:56 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 26th July 2023 10:08:14 am
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
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/backend/store-pro/pkg/store/store_api"
	"gitlab.volio.vn/tech/fountain/baselib/base"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/validor"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

const (
	KStoreIDKey        = "store_id"
	KStoreIDMigrateKey = "store_id_migrate"
)

func (ctrl *StoresController) CheckStoreIDExistence(ctx *fiber.Ctx) error {
	callback := ctx.Get("callback")

	if callback == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_DATA_INVALID), "callback is required!")
		v_log.V(3).Errorf("Middleware::CheckStoreIDExistence - callback error: %+v", err)

		return store_api.WriteError(ctx, err)
	}

	store := ctrl.GetStoresByAddress(callback)
	if store == nil {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_DATA_INVALID), "callback is invalid!")
		v_log.V(3).Errorf("Middleware::CheckStoreIDExistence - callback error: %+v", err)

		ctrl.dataMigratesCallback.MigrateStores()

		return store_api.WriteError(ctx, err)
	}
	ctx.Locals(KStoreIDMigrateKey, store.ID)

	storeMigrate := ctrl.dataMigratesCallback.GetByID(store.ID)
	ctx.Locals(KStoreIDKey, base.Int32ToString(storeMigrate.DataID))

	return ctx.Next()
}

func (ctrl *StoresController) registerCustomValidorRule() {
	// thêm -rule ở cuối để nhấn mạnh đây là rule tự tạo
	validor.NewValidator().RegisterRule("exist-store-rule", ctrl.CheckExistStore, "not found store for this request")
}

func (ctrl *StoresController) CheckExistStore(fl validator.FieldLevel) bool {
	storeID := fl.Field().String()
	return ctrl.dao.storesDAO.GetStore(storeID) != nil
}

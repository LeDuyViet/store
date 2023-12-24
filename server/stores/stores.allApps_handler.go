/* !!
 * File: stores.apps_handler.go
 * File Created: Thursday, 27th July 2023 3:51:18 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th July 2023 3:51:18 pm
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

package stores_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

func (api *StoresAPI) GetAllApps(ctx *fiber.Ctx) error {
	storeID := ctx.Params("id", "")
	if storeID == "" {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "store id is required")
		v_log.V(1).WithError(err).Errorf("StoreAPI::Delete - Error: %+v", err)

		return v_api.WriteError(ctx, err)
	}

	res := api.storesController.GetAllApps(storeID)

	return v_api.WriteSuccess(ctx, res)
}

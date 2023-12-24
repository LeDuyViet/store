/* !!
 * File: apps.go
 * File Created: Thursday, 27th July 2023 3:55:17 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th July 2023 3:55:17 pm
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
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
)

func (ctrl *StoresController) GetAllApps(storeID string) []*models.AppMD {
	return ctrl.appsCallback.GetAllByStore(storeID)
}

func (ctrl *StoresController) GetApps(storeID string, offset, limit int) []*models.AppMD {
	return ctrl.appsCallback.GetByStore(storeID, offset, limit)
}

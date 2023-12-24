/* !!
 * File: apps_impl.go
 * File Created: Wednesday, 26th July 2023 2:30:29 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 26th July 2023 2:30:29 pm
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

package apps_handler

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/apps"
	"gitlab.volio.vn/tech/backend/store-pro/biz/core/data_migrates"
	"gitlab.volio.vn/tech/fountain/baselib/v_net/v_api"
	"gitlab.volio.vn/tech/fountain/biz/core"
)

type AppsAPI struct {
	*v_api.VolioAPI

	appsController         *apps.AppsController
	dataMigratesController *data_migrates.DataMigratesController
}

func NewAppsAPI(controllers []core.CoreController) *AppsAPI {
	impl := &AppsAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *apps.AppsController:
			impl.appsController = x
		case *data_migrates.DataMigratesController:
			impl.dataMigratesController = x
		}
	}

	return impl
}

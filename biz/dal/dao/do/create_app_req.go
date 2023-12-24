/* !!
 * File: create_app_req.go
 * File Created: Thursday, 27th July 2023 4:41:24 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th July 2023 4:41:24 pm
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

package do

type CreateAppReq struct {
	ID string `db:"id" json:"id,omitempty"`

	StoreID     string `db:"store_id" json:"store_id,omitempty" validate:"required,uuid4,exist-store-rule"`
	Name        string `db:"name" json:"name,omitempty" validate:"required,gte=3,lte=128"`
	PackageName string `db:"package_name" json:"package_name,omitempty" validate:"required,gte=3,lte=128,unique-package-name-rule"`
	Photo       string `db:"photo" json:"photo,omitempty" validate:"url"`

	CustomFields []*CustomFieldTable `db:"-" json:"custom_fields"`
}

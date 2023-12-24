/* !!
 * File: custom_field_md.go
 * File Created: Monday, 10th July 2023 3:16:35 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd August 2023 2:47:46 pm
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

package models

// Các đối tượng như Category, Item có thể thêm mới custom field, custom field khi thêm mới có thể lựa chọn từ template có sẵn hoặc tạo mới hoàn toàn.
// Custom field template có scope giảm dần từ TEAM < USER < STORE < APP < CATEGORY
// Custom field data dùng để lưu trữ value cho các custom field

type CustomFieldMD struct {
	ID          string          `db:"id" json:"id,omitempty"`
	Name        string          `db:"name" json:"name,omitempty"`
	Status      CustomFieldType `db:"status" json:"status,omitempty"` // Private, Public
	Scope       string          `db:"scope" json:"scope,omitempty"`   // CATEGORY, ITEM
	ScopeID     string          `db:"scope_id" json:"scope_id,omitempty"`
	CreatedTime int32           `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32           `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}

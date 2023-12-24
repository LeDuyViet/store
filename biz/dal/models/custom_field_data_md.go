/* !!
 * File: custom_field_data_md.go
 * File Created: Monday, 10th July 2023 3:16:35 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd August 2023 3:02:28 pm
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

type CustomFieldDataMD struct {
	ID            string `db:"id" json:"id,omitempty"`
	CustomFieldID string `db:"custom_field_id" json:"custom_field_id,omitempty"`
	Value         string `db:"value" json:"value,omitempty"`
	CreatedTime   int32  `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime   int32  `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}

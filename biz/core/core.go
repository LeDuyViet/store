/* !!
 * File: core.go
 * File Created: Thursday, 3rd November 2022 4:07:06 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd November 2022 4:07:06 pm
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

package core

import (
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/do"
	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/models"
	"gitlab.volio.vn/tech/fountain/proto/v_proto"
)

// --------------------------------- CALLBACK -------------------------------

type StoresCallback interface {
	GetStore(id string) *models.StoreMD
	GetByPackageName(packageName string) []*models.StoreMD
	GetAll() []*models.StoreMD
}

type AppsCallback interface {
	GetApp(id string) *models.AppMD
	GetByStore(storeID string, offset, limit int) []*models.AppMD
	GetAllByStore(storeID string) []*models.AppMD
	CreateApp(app *do.CreateAppReq) *v_proto.VolioRpcError
}

type ModulesCallback interface {
	GetModule(id string) *models.ModuleMD
	CreateModule(module *models.ModuleMD) *v_proto.VolioRpcError
}

type CustomFieldsCallback interface {
	GetCustomField(id string) *models.CustomFieldMD
	CreateCustomField(customField *models.CustomFieldMD) *v_proto.VolioRpcError
	GetForPublicModule(moduleID, AppID string) []*do.StoreCustomFieldTableDO
	GetForModule(moduleID, AppID string) []*do.CustomFieldTable
	GetForPublicCategory(category_id, module_id string) []*do.StoreCustomFieldTableDO
	GetForPublicItem(category_id, module_id string) []*do.StoreCustomFieldTableDO
	GetData(id string) *models.CustomFieldDataMD
	GetManyData(customFields []*do.CustomFieldTable) []*models.CustomFieldDataMD
	GetByIDsData(ids []string) []*models.CustomFieldDataMD
	InsertManyData(md []*do.CustomFieldTable) *v_proto.VolioRpcError
	InsertData(customFieldTable *models.CustomFieldDataMD) *v_proto.VolioRpcError
	UpdateManyData(md []*do.CustomFieldTable) *v_proto.VolioRpcError
	GetByTableableData(CustomFieldTableableID, CustomFieldTableableType string) []*models.CustomFieldDataMD
	DeleteData(id ...string) *v_proto.VolioRpcError
}

type CategoriesCallback interface {
	GetCategory(id string) *models.CategoryMD
	CreateCategory(category *models.CategoryMD) *v_proto.VolioRpcError
	GetByApp(appID string, limit, offset int, isRoot ...bool) []*do.CategoriesDO
	GetAllByApp(appID string) []*models.CategoryMD
}

type ItemsCallback interface {
	GetItem(id string) *models.ItemMD
	CreateItem(item *models.ItemMD) *v_proto.VolioRpcError
}

type DataMigratesCallback interface {
	GetByDataMigrate(dataMigrate *models.DataMigratesMD) *models.DataMigratesMD
	CreateDataMigrate(dataMigrate *models.DataMigratesMD) *v_proto.VolioRpcError
	CheckExistInMigration(dataType string, dataId, storeId int32) (*models.DataMigratesMD, bool)
	GetByID(id string) *models.DataMigratesMD
	MigrateData(storeId, appId int32)
	MigrateStores()
}

// -------------------------------- Interface -------------------------------
type Admin interface {
	// AdminGetAll(parentID string, perPage, current_page int, path string) *do.ResPaginate
	AdminCreate(input interface{}) *v_proto.VolioRpcError
	AdminUpdate(input interface{}) *v_proto.VolioRpcError
	AdminGet(id string) interface{}
	AdminGets(limit, offset int) interface{}
	AdminDelete(id string) *v_proto.VolioRpcError
}

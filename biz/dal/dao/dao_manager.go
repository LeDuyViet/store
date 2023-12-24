/* !!
 * File: dao_manager.go
 * File Created: Thursday, 6th October 2022 3:05:27 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 6th October 2022 3:05:30 pm
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

package dao

import (
	"sync"

	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao/postgres_dao"
	"gitlab.volio.vn/tech/fountain/baselib/redis_client"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

const (
	STORES_CACHE = "stores_cache"

	STORES_DB_MASTER = "stores"
	STORES_DB_SLAVE  = "stores_slave"

	STORES_STORAGE = "stores_storage"
)

// ----------------------- PostgresSQL ------------------------

type PostgresDAOList struct {
	storesDAO            *postgres_dao.StoresDAO
	dataMigratesDAO      *postgres_dao.DataMigratesDAO
	appsDAO              *postgres_dao.AppsDAO
	modulesDAO           *postgres_dao.ModulesDAO
	customFieldsDAO      *postgres_dao.CustomFieldsDAO
	categoriesDAO        *postgres_dao.CategoriesDAO
	customFieldTablesDAO *postgres_dao.CustomFieldDataDAO
	itemsDAO             *postgres_dao.ItemsDAO
	regionsDAO           *postgres_dao.RegionsDAO
	appRegionsDAO        *postgres_dao.AppRegionsDAO
}

// PostgresDAOManager type
type PostgresDAOManager struct {
	daoListMap map[string]*PostgresDAOList
}

var postgresDAOManager = &PostgresDAOManager{make(map[string]*PostgresDAOList)}

// InstallMysqlDAOManager func
func InstallMysqlDAOManager(clients sync.Map) { /*map[string]*sql_client.SQLClient*/
	clients.Range(func(key, value interface{}) bool {
		k, _ := key.(string)
		v, _ := value.(*sql_client.SQLClient)

		daoList := &PostgresDAOList{}
		daoList.storesDAO = postgres_dao.NewStoresDAO(v.DB)
		daoList.dataMigratesDAO = postgres_dao.NewDataMigratesDAO(v.DB)
		daoList.appsDAO = postgres_dao.NewAppsDAO(v.DB)
		daoList.modulesDAO = postgres_dao.NewModulesDAO(v.DB)
		daoList.customFieldsDAO = postgres_dao.NewCustomFieldsDAO(v.DB)
		daoList.categoriesDAO = postgres_dao.NewCategoriesDAO(v.DB)
		daoList.customFieldTablesDAO = postgres_dao.NewCustomFieldDataDAO(v.DB)
		daoList.itemsDAO = postgres_dao.NewItemsDAO(v.DB)
		daoList.regionsDAO = postgres_dao.NewRegionsDAO(v.DB)
		daoList.appRegionsDAO = postgres_dao.NewAppRegionsDAO(v.DB)

		postgresDAOManager.daoListMap[k] = daoList
		return true
	})
}

// GetPostgresDAOListMap func
func GetPostgresDAOListMap() map[string]*PostgresDAOList {
	return postgresDAOManager.daoListMap
}

// GetPostgresDAOList func
func GetPostgresDAOList(dbName string) (daoList *PostgresDAOList) {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Error("GetPostgresDAOList - Not found daoList: ", dbName)
	}

	return
}

func GetStoresDAO(dbName string) *postgres_dao.StoresDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetStoresDAO - Not found daoList: %s", dbName)
	}

	return daoList.storesDAO
}

func GetDataMigratesDAO(dbName string) *postgres_dao.DataMigratesDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetDataMigratesDAO - Not found daoList: %s", dbName)
	}

	return daoList.dataMigratesDAO
}

func GetAppsDAO(dbName string) *postgres_dao.AppsDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetAppsDAO - Not found daoList: %s", dbName)
	}

	return daoList.appsDAO
}

func GetModulesDAO(dbName string) *postgres_dao.ModulesDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetModulesDAO - Not found daoList: %s", dbName)
	}

	return daoList.modulesDAO
}

func GetCustomFieldsDAO(dbName string) *postgres_dao.CustomFieldsDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetCustomFieldsDAO - Not found daoList: %s", dbName)
	}

	return daoList.customFieldsDAO
}

func GetCategoriesDAO(dbName string) *postgres_dao.CategoriesDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetCategoriesDAO - Not found daoList: %s", dbName)
	}

	return daoList.categoriesDAO
}

func GetCustomFieldDataDAO(dbName string) *postgres_dao.CustomFieldDataDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetCustomFieldDataDAO - Not found daoList: %s", dbName)
	}

	return daoList.customFieldTablesDAO
}

func GetItemsDAO(dbName string) *postgres_dao.ItemsDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetItemsDAO - Not found daoList: %s", dbName)
	}

	return daoList.itemsDAO
}

func GetRegionsDAO(dbName string) *postgres_dao.RegionsDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetRegionsDAO - Not found daoList: %s", dbName)
	}

	return daoList.regionsDAO
}

func GetAppRegionsDAO(dbName string) *postgres_dao.AppRegionsDAO {
	daoList, ok := postgresDAOManager.daoListMap[dbName]
	if !ok {
		v_log.V(1).Infof("GetRegionsDAO - Not found daoList: %s", dbName)
	}

	return daoList.appRegionsDAO
}

// ----------------------- Redis ------------------------
// RedisDAOList type
type RedisDAOList struct {
}

// RedisDAOManager type
type RedisDAOManager struct {
	daoListMap map[string]*RedisDAOList
}

var redisDAOManager = &RedisDAOManager{make(map[string]*RedisDAOList)}

// InstallRedisDAOManager type
func InstallRedisDAOManager(clients map[string]*redis_client.RedisPool) {
	for k, v := range clients {
		_ = v

		daoList := &RedisDAOList{}

		redisDAOManager.daoListMap[k] = daoList
	}
}

// GetRedisDAOList type
func GetRedisDAOList(redisName string) (daoList *RedisDAOList) {
	daoList, ok := redisDAOManager.daoListMap[redisName]
	if !ok {
		v_log.V(1).Infof("GetRedisDAOList - Not found daoList: %s", redisName)
	}
	return
}

// GetRedisDAOListMap type
func GetRedisDAOListMap() map[string]*RedisDAOList {
	return redisDAOManager.daoListMap
}

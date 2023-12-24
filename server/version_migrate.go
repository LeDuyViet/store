/* !!
 * File: version_migrate.go
 * File Created: Thursday, 3rd November 2022 4:09:59 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd November 2022 4:09:59 pm
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
	"time"

	"gitlab.volio.vn/tech/backend/store-pro/biz/dal/dao"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
)

   
// This is database
// Table custom_fields {
//     id varchar(36) [primary key]
//     name varchar
//     scope varchar // STORE, APP, CATEGORY, ITEM
//     scopeID varchar
//     status varchar // Public, Private
//     created_time timestamp
//     updated_time timestamp   
//     }
    
//     Table custom_field_values {
//     id varchar(36) [primary key]
//     custom_field_id varchar(36)
//     record_id varchar(36)
//     value varchar
//     // object varchar // STORE, APP, CATEGORY, ITEM
//     // objectID varchar
//     created_time timestamp
//     updated_time timestamp
//     }
    
//     Table custom_field_templates {
//     id varchar(36) [primary key]
//     name varchar
//     description varchar
//     scope varchar
//     creator varchar
//     type bool
//     created_time timestamp
//     updated_time timestamp
//     }
    
//     Table custom_field_template_fields {
//     id varchar(36) [primary key]
//     custom_field_template_id varchar(36)
//     custom_field_id varchar(36)
//     position int
//     created_time timestamp
//     updated_time timestamp
//     }
    
//     Ref: custom_field_template_fields.custom_field_template_id > custom_field_templates.id // many-to-one
    
//     Ref: custom_field_template_fields.custom_field_id > custom_fields.id // many-to-one
    
//     Table stores {
//       id varchar(36) [primary key]
//       name varchar
//       address varchar
//       is_active bool
//       access_key varchar
//       created_time integer
//       updated_time integer
//     }
    
//     Table apps {
//       id varchar(36) [primary key]
//       store_id varchar(36)
//       parent_id varchar(36)
//       name varchar
//       package_name varchar
//       thumbnail varchar
//       created_time integer
//       updated_time integer
//     }
    
//     Table categories {
//       id varchar(36) [primary key]
//       app_id varchar(36)
//       parent_id varchar(36)
//       // level smallint
//       priority smallint
//       name varchar
//       thumbnail varchar
//       photo varchar
//       status bool
//       is_pro bool
//       is_new bool
//       created_time integer
//       updated_time integer
//     }
    
//     Ref: categories.app_id > apps.id // many-to-one
    
//     Table items {
//       id varchar(36) [primary key]
//       category_id varchar(36)
//       priority smallint
//       name varchar
//       thumbnail varchar
//       photo varchar
//       status bool
//       is_pro bool
//       is_new bool
//       created_time integer
//       updated_time integer
//     }
    
//     Ref: items.category_id > categories.id // many-to-one
    
//     Table regions {
//       id varchar(36) [primary key]
//       name varchar
//       code varchar
//       icon varchar
//       is_active bool
//       created_time integer
//       updated_time integer
//     }
    
    
//     Table app_regions {
//       app_id varchar(36)
//       region_id varchar(36)
//       is_active bool
//       created_time integer
//       updated_time integer
//       primary key (app_id, region_id)
//     }
    
//     Ref: app_regions.app_id > apps.id // many-to-one
    
//     Ref: app_regions.region_id > regions.id // many-to-one
    
//     Table data_migrates {
//       store_id integer
//       data_id integer
//       data_type varchar(36)
//       migration_id varchar(36)
//       is_sync bool
//       created_time integer
//       updated_time integer
//       primary key (store_id, data_id, data_type)
//     }
    
//     Ref: data_migrates.store_id > stores.id // many-to-one
    
//     Table folders {
//       id varchar(36) [primary key]
//       name varchar
//       parent_id varchar
//       created_time integer
//       updated_time integer
//     }
    
//     Table files {
//       id varchar(36) [primary key]
//       name varchar
//       type varchar
//       size int
//       folder_id varchar(36)
//       tags varchar
//       created_time integer
//       updated_time integer
//     }
    
//     Ref: "folders"."parent_id" < "files"."id"

// DO NOT EDIT IT, ONLY ADDING MORE
func init() {
	timeNow := int32(time.Now().Unix())
	sql_client.InitVersionMigrate(dao.STORES_DB_MASTER)
	sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
		Name: "stores_tbl_init",
		Query: `CREATE TABLE public.stores
        (
            id varchar(36) NOT NULL,
            name VARCHAR(256) NULL,
            address VARCHAR(256) NOT NULL,
            is_active bool DEFAULT false NOT NULL,
            access_key VARCHAR(256) NULL,
            created_time integer DEFAULT 0,
            updated_time integer DEFAULT 0,
            PRIMARY KEY (id)
        );

        CREATE INDEX address_idx_on_stores_table ON stores (address);
        `,
		CreatedTime: timeNow,
	})

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "apps_tbl_init",
        Query: `
            CREATE TABLE public.apps (
                id varchar(36) NOT NULL,
                store_id VARCHAR(36) NOT NULL,
                name VARCHAR(256) NULL,
                package_name VARCHAR(256) NULL,
                thumbnail VARCHAR(256) NULL DEFAULT '',
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (id)
            );

            CREATE INDEX package_name_idx_on_apps_table ON apps (package_name);
        `,
        CreatedTime: timeNow,
    })

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "categories_tbl_init",
        Query: `
            CREATE TABLE public.categories (
                id VARCHAR(36) NOT NULL,
                parent_id VARCHAR(36) NOT NULL,
                app_id VARCHAR(36) NOT NULL,
                priority smallint null,
                name VARCHAR(256) NULL,
                thumbnail VARCHAR(256) NULL,
                photo VARCHAR(256) NULL,
                status bool DEFAULT false NOT NULL,
                is_pro bool DEFAULT false NOT NULL,
                is_new bool DEFAULT false NOT NULL,
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (id)
            );
        `,
        CreatedTime: timeNow,
    })

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "items_tbl_init",
        Query: `
            CREATE TABLE public.items (
                id VARCHAR(36) NOT NULL,
                category_id VARCHAR(36) NOT NULL,
                priority smallint null,
                name VARCHAR(500) NULL,
                thumbnail VARCHAR(256) NULL,
                icon VARCHAR(256) NULL,
                status bool DEFAULT false NOT NULL,
                is_pro bool DEFAULT false NOT NULL,
                is_new bool DEFAULT false NOT NULL,
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (id)
            );
        `,
        CreatedTime: timeNow,
    })


    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "custom_fields_tbl_init",
        Query: `
            CREATE TABLE public.custom_fields (
                id VARCHAR(36) NOT NULL,
                name VARCHAR(256) NULL,
                scope VARCHAR(256) NULL,
                scopeID VARCHAR(36) NULL,
                type smallint null,
                status bool DEFAULT false NOT NULL,
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (id)
            );
        `,
        CreatedTime: timeNow,
    })


    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "custom_field_data_tbl_init",
        Query: `
        CREATE TABLE public.custom_field_data (
            id VARCHAR(36) NOT NULL,
            custom_field_id VARCHAR(36) NOT NULL,
            record_id VARCHAR(36) NOT NULL,
            value VARCHAR(256) NULL,
            created_time integer DEFAULT 0,
            updated_time integer DEFAULT 0,
            `,
        CreatedTime: timeNow,
    })

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "regions_tbl_init",
        Query: `
            CREATE TABLE public.regions (
                id VARCHAR(36) NOT NULL,
                name VARCHAR(256) NULL,
                code VARCHAR(256) NULL,
                icon VARCHAR(256) NULL,
                is_active bool DEFAULT false NOT NULL,
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (id)
            );
        `,
        CreatedTime: timeNow,
    })

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "app_regions_tbl_init",
        Query: `
            CREATE TABLE public.app_regions (
                app_id VARCHAR(36) NULL,
                region_id VARCHAR(36) NULL,
                is_active bool DEFAULT false NOT NULL,
                created_time integer DEFAULT 0,
                updated_time integer DEFAULT 0,
                PRIMARY KEY (app_id, region_id)
                );
            `,
        CreatedTime: timeNow,
    })
            
    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
        Name: "data_migrate_tbl_init",
        Query: `
        CREATE TYPE DATA_MIGRATE_TYPE AS ENUM('STORE', 'APP', 'MODULE', 'CATEGORY', 'ITEM', 'CUSTOM_FIELD', 'CUSTOM_FIELD_TABLE');
        
        CREATE TABLE public.data_migrates (
            store_id integer NOT NULL,
            data_id integer NOT NULL,
            data_type DATA_MIGRATE_TYPE NOT NULL,
            migration_id VARCHAR(36) NOT NULL,
            is_sync bool DEFAULT false NOT NULL,
            created_time integer DEFAULT 0,
            updated_time integer DEFAULT 0,
            PRIMARY KEY (store_id, data_id, data_type)
        );
        
        CREATE INDEX migration_idx_on_data_migrates_table ON data_migrates (store_id, data_id, data_type);
        `,
        CreatedTime: timeNow,
    })

    sql_client.AddVersionMigrateQuery(dao.STORES_DB_MASTER, &sql_client.VersionModels{
            Name: "set_UNIQUE_name_and_code_in_regions_table",
            Query: `
            ALTER TABLE public.regions
                ALTER COLUMN name SET NOT NULL,
                ADD UNIQUE(name),
                ALTER COLUMN code SET NOT NULL,
                ADD UNIQUE(code)
            `,
        CreatedTime: timeNow,
    })
}

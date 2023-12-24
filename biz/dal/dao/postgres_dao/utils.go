/* !!
 * File: utils.go
 * File Created: Friday, 13th January 2023 10:10:54 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Friday, 13th January 2023 10:10:54 am
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

package postgres_dao

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.volio.vn/tech/fountain/baselib/sql_client"
	"gitlab.volio.vn/tech/fountain/baselib/v_log"
)

func queryDataParser[T any](conn *sqlx.DB, query string, fnAfterParse func(*T), args ...interface{}) (*T, error) {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	rows, err := conn.QueryxContext(ctx, query, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("parseListData - Error: %+v", err)

		return nil, err
	}

	do := new(T)
	if rows.Next() {
		if err := rows.StructScan(do); err != nil {
			v_log.V(1).WithError(err).Errorf("parseData - Error: %+v", err)

			return do, err
		}
	} else {
		return nil, fmt.Errorf("not found resource")
	}

	if err := rows.Err(); err != nil {
		v_log.V(1).WithError(err).Errorf("parseData- Error: %+v", err)

		return nil, err
	}

	if fnAfterParse != nil {
		fnAfterParse(do)
	}

	return do, nil
}

func queryListDataParser[T any](conn *sqlx.DB, query string, fnAfterParse func(*T), args ...interface{}) ([]*T, error) {
	ctx, cancel := sql_client.CreateDefaultCtx()
	defer cancel()

	rows, err := conn.QueryxContext(ctx, query, args...)
	if err != nil {
		v_log.V(1).WithError(err).Errorf("parseListData - Error: %+v", err)

		return nil, err
	}

	res := make([]*T, 0)

	for rows.Next() {
		do := new(T)
		if err := rows.StructScan(do); err == nil {
			if fnAfterParse != nil {
				fnAfterParse(do)
			}

			res = append(res, do)
		} else {
			v_log.V(1).WithError(err).Errorf("parseListData - Error: %+v", err)
		}
	}

	if err := rows.Err(); err != nil {
		v_log.V(1).WithError(err).Errorf("parseListData - Error: %+v", err)

		return nil, err
	}

	return res, nil
}
func countTotal(conn *sqlx.DB, table_name, conditions string, args ...interface{}) int {
	total := 0
	var err error

	if conditions != "" {
		query := fmt.Sprintf("SELECT COUNT(1) FROM %s where %s", table_name, conditions)
		err = conn.Get(&total, query, args...)
	} else {
		query := fmt.Sprintf("SELECT COUNT(1) FROM %s", table_name)
		err = conn.Get(&total, query, args...)
	}

	if err != nil {
		v_log.V(1).WithError(err).Errorf("CountTotal - Error: %+v", err)
		return total
	}
	return total
}

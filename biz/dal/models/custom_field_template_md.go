/* !!
 * File: custom_field_template_md.go
 * File Created: Thursday, 3rd August 2023 2:46:21 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 3rd August 2023 2:47:40 pm
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

type CustomFieldType string

const (
	KCustomFieldTypeDropDown = "DROPDOWN"      // Use dropdowns to give consistent options - even use colors!
	KCustomFieldTypeRating   = "RATING"        // Use emojis to rate or rank
	KCustomFieldTypeMoney    = "MONEY"         // Add information such as budget, cost, or price.
	KCustomFieldTypeFormula  = "FORMULA"       // Calculate fields using formulas and more.
	KCustomFieldTypeAreaText = "AREA_TEXT"     // Long Text - Capture lots of text for things like notes, descriptions, addresses, or anything that takes up more than one line.
	KCustomFieldTypeNumber   = "NUMBER"        // Use number fields for accounting, inventory, or tracking.
	KCustomFieldTypeLabels   = "LABELS"        // Tags - Add one or more labels with colors
	KCustomFieldTypePeople   = "PEOPLE"        // Select people in your Workspace
	KCustomFieldText         = "TEXT"          // Short Text -Capture short text for things like names, locations, or anything you want in just one line.
	KCustomFieldDate         = "DATE"          // Add any date to your task.
	KCustomFieldWebsite      = "WEBSITE"       // Add websites that are associated with the task.
	KCustomFieldCheckbox     = "CHECKBOX"      // Yes or no? Add a simple true or false checkbox.
	KCustomFieldEmail        = "EMAIL"         // Track clients, vendors, leads, and more by entering emails.
	KCustomFieldPhone        = "PHONE"         // Use Store as a CRM by adding phone numbers.
	KCustomFieldFiles        = "FILES"         // Add one or more files to your object.
	KCustomFieldLocation     = "LOCATION"      // Add an address or a place to your task.
	KCustomFieldProgress     = "PROGRESS"      // Manually track progress of anything
	KCustomFieldProgressAuto = "PROGRESS_AUTO" // Automatically track completion of object
	KCustomFieldObject       = "OBJECT"        // Link to other Store Object
)

type CustomFieldTemplateMD struct {
	ID          string          `db:"id" json:"id,omitempty"`
	Name        string          `db:"name" json:"name,omitempty"`
	Type        CustomFieldType `db:"type" json:"type,omitempty"`
	ScopeID     string          `db:"scope_id" json:"scope_id,omitempty"`
	Scope       bool            `db:"scope" json:"scope,omitempty"`               // TEAM, USER, STORE, APP, CATEGORY
	CreatedTime int32           `db:"created_time" json:"created_time,omitempty"` // seconds UTC
	UpdatedTime int32           `db:"updated_time" json:"updated_time,omitempty"` // seconds UTC
}

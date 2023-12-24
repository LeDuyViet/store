/* !!
 * File: main.go
 * File Created: Thursday, 27th July 2023 11:57:10 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th July 2023 11:57:10 am
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

package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Hits struct {
	MaxScore float64 `json:"max_score"`
	HitsList []Hit   `json:"hits"`
}

type Hit struct {
	Index  string              `json:"_index"`
	ID     string              `json:"_id"`
	Score  float64             `json:"_score"`
	Fields map[string][]string `json:"fields"`
}

type PathInfo struct {
	Path       string
	Count      int
	Percentage float64
}

func main() {
	// Đọc nội dung từ tệp tin "data.json"
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Lỗi khi đọc tệp tin:", err)
		return
	}

	// Parse dữ liệu JSON vào struct
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Lỗi khi giải mã dữ liệu JSON:", err)
		return
	}

	// Lấy danh sách các hits từ dữ liệu JSON
	hitsData, ok := result["hits"].(map[string]interface{})
	if !ok {
		fmt.Println("Không tìm thấy danh sách hits trong kết quả")
		return
	}

	hitsList, ok := hitsData["hits"].([]interface{})
	if !ok {
		fmt.Println("Không tìm thấy danh sách hits trong kết quả")
		return
	}

	// Loại bỏ truy vấn trong path và thống kê tần số các path cơ bản
	pathCounts := make(map[string]int)
	totalPaths := 0

	for _, hitData := range hitsList {
		hit, ok := hitData.(map[string]interface{})
		if !ok {
			fmt.Println("Không thể lấy thông tin hit từ dữ liệu JSON")
			continue
		}

		fieldsData, ok := hit["fields"].(map[string]interface{})
		if !ok {
			fmt.Println("Không thể lấy thông tin fields từ dữ liệu JSON")
			continue
		}

		methodValues, ok := fieldsData["method.keyword"].([]interface{})
		if !ok || len(methodValues) == 0 {
			fmt.Println("Không tìm thấy giá trị method.keyword trong dữ liệu JSON")
			continue
		}

		pathValues, ok := fieldsData["path.keyword"].([]interface{})
		if !ok || len(pathValues) == 0 {
			fmt.Println("Không tìm thấy giá trị path.keyword trong dữ liệu JSON")
			continue
		}

		// Loại bỏ truy vấn trong path và thống kê tần số các path kết hợp với method
		method, ok := methodValues[0].(string)
		if !ok {
			fmt.Println("Giá trị method không hợp lệ")
			continue
		}

		path, ok := pathValues[0].(string)
		if !ok {
			fmt.Println("Giá trị path không hợp lệ")
			continue
		}

		u, err := url.Parse(path)
		if err != nil {
			fmt.Println("Lỗi khi phân tích URL:", err)
			continue
		}

		re := regexp.MustCompile(`modules/(\d+)/categories`)

		u.RawQuery = ""
		cleanPath := u.String()
		key := fmt.Sprintf("%s - %s", method, strings.TrimPrefix(cleanPath, "/myzgos"))

		if matches := re.FindStringSubmatch(key); len(matches) == 2 {
			key = re.ReplaceAllString(key, "modules/:module_id/categories")
		}

		u.RawQuery = ""

		pathCounts[key]++
		totalPaths++
	}

	// Tạo slice để lưu trữ thông tin về các path và số lượng tần suất của chúng
	var pathInfoList []PathInfo

	// Thêm thông tin path và số lượng tần suất vào slice
	for path, count := range pathCounts {
		percentage := float64(count) / float64(totalPaths) * 100
		pathInfoList = append(pathInfoList, PathInfo{Path: path, Count: count, Percentage: percentage})
	}

	// Sắp xếp slice theo số lượng giảm dần
	sort.Slice(pathInfoList, func(i, j int) bool {
		return pathInfoList[i].Count > pathInfoList[j].Count
	})

	// In kết quả thống kê tần số của mỗi path kết hợp với method
	fmt.Printf("Thống kê tần số của mỗi path (%d bản ghi):\n", totalPaths)
	for _, pathInfo := range pathInfoList {
		fmt.Printf("%s: %d (Tỉ lệ phần trăm: %.2f%%)\n", pathInfo.Path, pathInfo.Count, pathInfo.Percentage)
	}
}

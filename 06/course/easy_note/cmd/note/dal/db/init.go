// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true, //缓存起来  查询更快
			SkipDefaultTransaction: true, //当使用gorm的hook和关联创建，使用false 保持数据一致性，其他可以为ture
		},
	)
	if err != nil {
		panic(err)
	}
	// gorm使用opentracing插件
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	//迁移表，如果存在就返回 ，不存在就报恐慌
	m := DB.Migrator()
	if m.HasTable(&Note{}) {
		return
	}
	if err = m.CreateTable(&Note{}); err != nil {
		panic(err)
	}
}

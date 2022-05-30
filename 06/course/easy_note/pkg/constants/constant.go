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

package constants

const (
	NoteTableName           = "note"       // 笔记表名
	UserTableName           = "user"       // 用户表名
	SecretKey               = "secret key" // jwt 变量
	IdentityKey             = "id"         //同上
	Total                   = "total"      // gorm 查询用到的变量
	Notes                   = "notes"      //同上
	NoteID                  = "note_id"    //同上
	ApiServiceName          = "demoapi"    //三类服务
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local" // 数据源名
	EtcdAddress             = "127.0.0.1:2379"                                                           //etcd 地址
	CPURateLimit    float64 = 80.0                                                                       //cpu 限制
	DefaultLimit            = 10                                                                         //默认分页限制
)

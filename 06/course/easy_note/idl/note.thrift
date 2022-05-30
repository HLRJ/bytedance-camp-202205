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
// 命名空间为notedemo
namespace go notedemo
// 每个响应都会返回BaseResp
struct BaseResp {
    1:i64 status_code  // 状态码
    2:string status_message  // 状态消息
    3:i64 service_time  // 服务时间
}
// 笔记的结构体
struct Note {
    1:i64 note_id
    2:i64 user_id
    3:string user_name
    4:string user_avatar  //用户头像  为了简化  设置成了string 并没有设置成url
    5:string title  //笔记标题
    6:string content //笔记内容
    7:i64 create_time //笔记创建时间
}
// 创建笔记的请求
struct CreateNoteRequest {
    1:string title  // 笔记的标题
    2:string content // 内容
    3:i64 user_id //用户的id
}
// 创建笔记的响应
struct CreateNoteResponse {
    1:BaseResp base_resp
}
// 删除笔记请求
struct DeleteNoteRequest {
    1:i64 note_id
    2:i64 user_id
}
// 响应
struct DeleteNoteResponse {
    1:BaseResp base_resp
}
// 更新笔记的请求
struct UpdateNoteRequest {
    1:i64 note_id
    2:i64 user_id
    3:optional string title //可选的字段
    4:optional string content //可选的字段  不传值的话 即为 空字符串
}

struct UpdateNoteResponse {
    1:BaseResp base_resp
}
// 批量获取笔记的请求
struct MGetNoteRequest {
    1:list<i64> note_ids  //传递一组笔记id
}
//
struct MGetNoteResponse {
    1:list<Note> notes
    2:BaseResp base_resp
}
// 查询笔记请求
struct QueryNoteRequest {
    1:i64 user_id //用户id
    2:optional string search_key  //使用关键字来模糊搜索
    3:i64 offset  //用offset 和 limit 来模糊搜索
    4:i64 limit
}
// 返回查询的结果，包含一组笔记 笔记总的数量和返回状态
struct QueryNoteResponse {
    1:list<Note> notes
    2:i64 total
    3:BaseResp base_resp
}
// 五个服务
service NoteService {
    CreateNoteResponse CreateNote(1:CreateNoteRequest req)
    MGetNoteResponse MGetNote(1:MGetNoteRequest req)
    DeleteNoteResponse DeleteNote(1:DeleteNoteRequest req)
    QueryNoteResponse QueryNote(1:QueryNoteRequest req)
    UpdateNoteResponse UpdateNote(1:UpdateNoteRequest req)
}
// 使用ubuntu20.04生成时,命令如下
//kitex -module github.com/cloudwego/kitex-examples/bizdemo/easy_note -service note note.thrift

// Copyright 2018 JDCLOUD.COM
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
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type RoleInfo struct {

    /* 角色ID (Optional) */
    RoleId string `json:"roleId"`

    /* 角色名称 (Optional) */
    RoleName string `json:"roleName"`

    /* 角色类型，2-服务相关角色，3-服务角色，4-用户角色 (Optional) */
    Type int `json:"type"`

    /* 信任实体信息 (Optional) */
    AssumeRolePolicyDocument string `json:"assumeRolePolicyDocument"`

    /* 描述，0~256个字符 (Optional) */
    Description string `json:"description"`

    /* 最大会话时长3600~43200秒，默认3600秒 (Optional) */
    MaxSessionDuration int `json:"maxSessionDuration"`

    /* 京东云资源标识(jrn) (Optional) */
    Jrn string `json:"jrn"`

    /* 创建角色的时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 角色所属主账号 (Optional) */
    Account string `json:"account"`
}

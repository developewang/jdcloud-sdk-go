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


type IotcoreInstance struct {

    /* 实例标识 (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 管理平台地址 (Optional) */
    PortalAddress string `json:"portalAddress"`

    /* 设备连接域名 (Optional) */
    LinkDomain string `json:"linkDomain"`

    /* 实例状态
创建中：CREATING
运行中：RUNNING
启动中：STARTING
欠费停服：ARREARAGE_STOP
过期停服：EXPIRE_STOP
异常：ERROR    
 (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 创建进度，取值范围：[0,100] (Optional) */
    Progress int `json:"progress"`

    /* 创建时间，Unix时间戳，精确到毫秒，时间为东八区（UTC/GMT+08:00） (Optional) */
    CreatedTime int64 `json:"createdTime"`

    /* 更新时间，Unix时间戳，精确到毫秒，时间为东八区（UTC/GMT+08:00） (Optional) */
    UpdatedTime int64 `json:"updatedTime"`

    /* 到期时间，Unix时间戳，精确到毫秒，时间为东八区（UTC/GMT+08:00） (Optional) */
    ExpiredTime int64 `json:"expiredTime"`
}

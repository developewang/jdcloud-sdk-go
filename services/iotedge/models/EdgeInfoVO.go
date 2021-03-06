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


type EdgeInfoVO struct {

    /* 系统流水号 (Optional) */
    Uuid string `json:"uuid"`

    /* Edge的唯一编号 (Optional) */
    EdgeId string `json:"edgeId"`

    /* Edge的名称 (Optional) */
    EdgeName string `json:"edgeName"`

    /* 是否在线【0-离线，1-在线】 (Optional) */
    EdgeStatus int `json:"edgeStatus"`

    /* 边缘计算说明 (Optional) */
    EdgeDesc string `json:"edgeDesc"`

    /* Edge版本 (Optional) */
    EdgeVersion string `json:"edgeVersion"`

    /* Edge创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 最后在线时间 (Optional) */
    LastOnlineTime string `json:"lastOnlineTime"`

    /* 最后开机时间 (Optional) */
    LastTurnOnTime string `json:"lastTurnOnTime"`

    /* IoT Hub实例编号 (Optional) */
    IothubInstanceId string `json:"iothubInstanceId"`

    /* IoT Hub实例名称 (Optional) */
    IothubInstanceName string `json:"iothubInstanceName"`
}

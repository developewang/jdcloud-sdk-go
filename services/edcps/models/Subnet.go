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


type Subnet struct {

    /* 地域代码, 如cn-east-tz1 (Optional) */
    Region string `json:"region"`

    /* 可用区, 如cn-east-tz1a (Optional) */
    Az string `json:"az"`

    /* 子网ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 子网名称 (Optional) */
    Name string `json:"name"`

    /* 子网CIDR (Optional) */
    Cidr string `json:"cidr"`

    /* 私有网络Id (Optional) */
    VpcId string `json:"vpcId"`

    /* 私有网络名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* 可用ip数量 (Optional) */
    AvailableIpCount int `json:"availableIpCount"`

    /* 总ip数量 (Optional) */
    TotalIpCount int `json:"totalIpCount"`

    /* 网络类型 (Optional) */
    NetworkType string `json:"networkType"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}

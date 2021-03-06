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


type ProtocolProps struct {

    /* 协议属性名称 (Optional) */
    PropName string `json:"propName"`

    /* 协议属性描述 (Optional) */
    PropDesc string `json:"propDesc"`

    /* 协议属性值最大长度 (Optional) */
    MaxLength string `json:"maxLength"`

    /* 协议属性默认值 (Optional) */
    DefaultValue string `json:"defaultValue"`

    /* 协议属性是否必填，0-非必填, 1-必填 (Optional) */
    Required string `json:"required"`
}

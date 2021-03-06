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


type StackOut struct {

    /* 资源栈运行操作 (Optional) */
    Action string `json:"action"`

    /* 资源栈能否更新 (Optional) */
    CanUpdate bool `json:"canUpdate"`

    /* 资源栈创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 资源栈是否回滚设置 (Optional) */
    DisableRollback bool `json:"disableRollback"`

    /* 资源栈ID (Optional) */
    Id string `json:"id"`

    /* 资源栈输入信息，JSON格式，此字段只在查询资源栈详情describeStack时有值 (Optional) */
    Input interface{} `json:"input"`

    /* 资源栈名称 (Optional) */
    Name string `json:"name"`

    /* 资源栈输出信息，JSON格式, 此字段只在查询资源栈详情describeStack时有值 (Optional) */
    Output interface{} `json:"output"`

    /* 资源栈所在区域 (Optional) */
    Region string `json:"region"`

    /* 资源栈版本 (Optional) */
    StackVersion int64 `json:"stackVersion"`

    /* 资源栈运行状态 (Optional) */
    Status string `json:"status"`

    /* 资源栈运行状态原因 (Optional) */
    StatusReason string `json:"statusReason"`

    /* 资源栈使用的template ID (Optional) */
    TemplateId string `json:"templateId"`

    /* 资源栈超时时间 (Optional) */
    Timeout int64 `json:"timeout"`

    /* 资源栈更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}

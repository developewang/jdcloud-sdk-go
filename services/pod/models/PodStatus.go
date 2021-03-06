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


type PodStatus struct {

    /* pod当前状态 (Optional) */
    Phase string `json:"phase"`

    /* （简要）pod处于当前状态的原因 (Optional) */
    Reason string `json:"reason"`

    /* pod处于当前状态原因的详细信息 (Optional) */
    Message string `json:"message"`

    /* 分配给pod的IP地址。至少在集群内是可路由的。未分配则为空。 (Optional) */
    PodIP string `json:"podIP"`

    /* 目前pod的状态。 (Optional) */
    Conditions []PodCondition `json:"conditions"`

    /* Pod生命周期开始的时间。 (Optional) */
    StartTime string `json:"startTime"`
}

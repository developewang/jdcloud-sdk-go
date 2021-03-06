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


type CreateSgRule struct {

    /* 规则类型，ingress、egress  */
    RuleType string `json:"ruleType"`

    /* 协议，tcp、udp、icmp 或者 all  */
    Protocol string `json:"protocol"`

    /* 起始端口  */
    FromPort int `json:"fromPort"`

    /* 终止端口  */
    ToPort int `json:"toPort"`

    /* 安全组ID  */
    SecurityGroupId string `json:"securityGroupId"`

    /* 网络类型，internet、intranet (Optional) */
    NicType *string `json:"nicType"`

    /* 认证策略，accept、drop (Optional) */
    Policy *string `json:"policy"`

    /* 认证策略的权重，1-100。 (Optional) */
    Priority *int `json:"priority"`

    /* 目标IP地址范围  */
    CidrIp string `json:"cidrIp"`
}

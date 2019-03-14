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


type PasswordPolicy struct {

    /* 密码长度，6-20之间  */
    Length int `json:"length"`

    /* 密码有效期，0-1095之间  */
    Age int `json:"age"`

    /* 过期重置类型：0-主账号重置，1-子账号重置  */
    ExpirationOperation int `json:"expirationOperation"`

    /* 历史密码检查次数,0-10之间  */
    ReusePrevention int `json:"reusePrevention"`

    /* 密码重试次数,1-16之间  */
    RetryTimes int `json:"retryTimes"`

    /*   */
    ValidLoginDuration int `json:"validLoginDuration"`

    /* 密码字符类型，至少包含一种  */
    Rule PasswordPolicyRule `json:"rule"`
}

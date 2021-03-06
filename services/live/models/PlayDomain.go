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


type PlayDomain struct {

    /* 播放域名 (Optional) */
    PlayDomain string `json:"playDomain"`

    /* 播放域名Cname (Optional) */
    PlayDomainCname string `json:"playDomainCname"`

    /* 直播域名状态
  online: 启用
  offline: 停用
  configuring: 配置中
  configure_failed: 配置失败
  checking: 正在审核
  check_failed: 审核失败
 (Optional) */
    DomainStatus string `json:"domainStatus"`

    /* 播放域名类型
  normal: 普通播放域名
  restart: 回看域名
 (Optional) */
    PlayType string `json:"playType"`

    /* 域名创建时间
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
 (Optional) */
    CreateTime string `json:"createTime"`

    /* 域名更新时间
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
 (Optional) */
    UpdateTime string `json:"updateTime"`
}

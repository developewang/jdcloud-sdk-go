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


type EsLogEvent struct {

    /* 产生时间 (Optional) */
    AccessTime int `json:"accessTime"`

    /* 源ip (Optional) */
    RemoteAddr string `json:"remoteAddr"`

    /* 域名 (Optional) */
    Domain string `json:"domain"`

    /* 来源地区 (Optional) */
    Area string `json:"area"`

    /* 方法 (Optional) */
    Method string `json:"method"`

    /* url (Optional) */
    Url string `json:"url"`

    /* 恶意负载 (Optional) */
    PayLoad string `json:"payLoad"`

    /* 响应信息 (Optional) */
    Status string `json:"status"`

    /* 日志类型 (Optional) */
    LogType string `json:"logType"`

    /* 动作 (Optional) */
    Action string `json:"action"`

    /* 请求ID (Optional) */
    RequestId string `json:"requestId"`

    /* 回源错误信息 (Optional) */
    UpstreamErr string `json:"upstreamErr"`
}

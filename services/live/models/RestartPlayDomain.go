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


type RestartPlayDomain struct {

    /* 直播播放域名 (Optional) */
    PlayDomain string `json:"playDomain"`

    /* 直播时移状态:
  - on表示开启
  - off表示关闭
 (Optional) */
    RestartStatus string `json:"restartStatus"`

    /* 播放域名类型:
  - normal  一般的播放域名
  - restart 回看播放域名 (Optional) */
    PlayType string `json:"playType"`
}

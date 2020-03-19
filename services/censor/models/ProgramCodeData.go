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


type ProgramCodeData struct {

    /* 以图片左上角为坐标原点，小程序码区域左上角到y轴距离 (Optional) */
    X int `json:"x"`

    /* 以图片左上角为坐标原点，小程序码区域左上角到x轴距离 (Optional) */
    Y int `json:"y"`

    /* 小程序码区域宽度 (Optional) */
    W int `json:"w"`

    /* 小程序码区域高度 (Optional) */
    H int `json:"h"`
}

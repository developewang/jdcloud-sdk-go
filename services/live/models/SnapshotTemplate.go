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


type SnapshotTemplate struct {

    /* 截图格式 (Optional) */
    Format string `json:"format"`

    /* 宽 (Optional) */
    Width int `json:"width"`

    /* 高 (Optional) */
    Height int `json:"height"`

    /* 文件类型 (Optional) */
    FillType int `json:"fillType"`

    /* 截图 (Optional) */
    SnapshotInterval int `json:"snapshotInterval"`

    /* 保存模式 (Optional) */
    SaveMode int `json:"saveMode"`

    /* 存储桶 (Optional) */
    SaveBucket string `json:"saveBucket"`

    /* 存储endPoint (Optional) */
    SaveEndpoint string `json:"saveEndpoint"`

    /* 录制模板自定义名称 (Optional) */
    Template string `json:"template"`
}

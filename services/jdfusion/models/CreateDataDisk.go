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


type CreateDataDisk struct {

    /* 云注册信息ID (Optional) */
    CloudID *string `json:"cloudID"`

    /* 云硬盘ID (Optional) */
    Id *string `json:"id"`

    /* 云硬盘名称 (Optional) */
    Name *string `json:"name"`

    /* 云硬盘描述 (Optional) */
    Description *string `json:"description"`

    /* 磁盘大小,单位为 GiB  */
    DiskSizeGB int `json:"diskSizeGB"`

    /* 挂载信息 (Optional) */
    Attachments []ArrayDiskAttachment `json:"attachments"`

    /* 云硬盘所属AZ  */
    Az string `json:"az"`

    /* 磁盘类型,取值为 ssd 或 premium-hdd (Optional) */
    DiskType *string `json:"diskType"`

    /* 创建该云硬盘的快照ID (Optional) */
    SnapshotId *string `json:"snapshotId"`

    /* 云硬盘状态 (Optional) */
    Status *string `json:"status"`

    /*  (Optional) */
    Tags []Tag `json:"tags"`

    /* 创建时间 (Optional) */
    CreatedTime *string `json:"createdTime"`
}

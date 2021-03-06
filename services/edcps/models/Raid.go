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


type Raid struct {

    /* 磁盘类型, 如 system/data (Optional) */
    VolumeType string `json:"volumeType"`

    /* 设备详情 (Optional) */
    VolumeDetail string `json:"volumeDetail"`

    /* RAID类型ID (Optional) */
    RaidTypeId string `json:"raidTypeId"`

    /* RAID类型, 如 NORAID, RAID0, RAID1 (Optional) */
    RaidType string `json:"raidType"`

    /* 云物理服务器类型, 如 edcps.c.normal1 (Optional) */
    DeviceType string `json:"deviceType"`

    /* RAID类型描述 (Optional) */
    Description string `json:"description"`
}

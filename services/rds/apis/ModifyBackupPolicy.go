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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type ModifyBackupPolicyRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00 (Optional) */
    StartWindow *string `json:"startWindow"`

    /* binlog本地保留周期，单位小时,范围24-168 (Optional) */
    BinlogRetentionPeriod *int `json:"binlogRetentionPeriod"`

    /* binlog本地占用空间上限，单位%，范围1-50 (Optional) */
    BinlogUsageLimit *int `json:"binlogUsageLimit"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyBackupPolicyRequest(
    regionId string,
    instanceId string,
) *ModifyBackupPolicyRequest {

	return &ModifyBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyBackupPolicy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startWindow: 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00 (Optional)
 * param binlogRetentionPeriod: binlog本地保留周期，单位小时,范围24-168 (Optional)
 * param binlogUsageLimit: binlog本地占用空间上限，单位%，范围1-50 (Optional)
 */
func NewModifyBackupPolicyRequestWithAllParams(
    regionId string,
    instanceId string,
    startWindow *string,
    binlogRetentionPeriod *int,
    binlogUsageLimit *int,
) *ModifyBackupPolicyRequest {

    return &ModifyBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyBackupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        StartWindow: startWindow,
        BinlogRetentionPeriod: binlogRetentionPeriod,
        BinlogUsageLimit: binlogUsageLimit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyBackupPolicyRequestWithoutParam() *ModifyBackupPolicyRequest {

    return &ModifyBackupPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyBackupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyBackupPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *ModifyBackupPolicyRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param startWindow: 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00(Optional) */
func (r *ModifyBackupPolicyRequest) SetStartWindow(startWindow string) {
    r.StartWindow = &startWindow
}

/* param binlogRetentionPeriod: binlog本地保留周期，单位小时,范围24-168(Optional) */
func (r *ModifyBackupPolicyRequest) SetBinlogRetentionPeriod(binlogRetentionPeriod int) {
    r.BinlogRetentionPeriod = &binlogRetentionPeriod
}

/* param binlogUsageLimit: binlog本地占用空间上限，单位%，范围1-50(Optional) */
func (r *ModifyBackupPolicyRequest) SetBinlogUsageLimit(binlogUsageLimit int) {
    r.BinlogUsageLimit = &binlogUsageLimit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyBackupPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyBackupPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyBackupPolicyResult `json:"result"`
}

type ModifyBackupPolicyResult struct {
}
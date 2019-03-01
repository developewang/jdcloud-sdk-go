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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeProbeHistoryRequest struct {

    core.JDCloudRequest

    /* 可用性监控task_id, id长度(0,50]  */
    ProbeTaskID string `json:"probeTaskID"`

    /* 探测源id，  id长度（0,50]  */
    ProbeID string `json:"probeID"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间往前三天，早于3d时，将被重置为3d） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，默认为当前时间） (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param probeTaskID: 可用性监控task_id, id长度(0,50] (Required)
 * param probeID: 探测源id，  id长度（0,50] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProbeHistoryRequest(
    probeTaskID string,
    probeID string,
) *DescribeProbeHistoryRequest {

	return &DescribeProbeHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/am/probeTask/{probeTaskID}/probe/{probeID}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ProbeTaskID: probeTaskID,
        ProbeID: probeID,
	}
}

/*
 * param probeTaskID: 可用性监控task_id, id长度(0,50] (Required)
 * param probeID: 探测源id，  id长度（0,50] (Required)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间往前三天，早于3d时，将被重置为3d） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，默认为当前时间） (Optional)
 */
func NewDescribeProbeHistoryRequestWithAllParams(
    probeTaskID string,
    probeID string,
    startTime *string,
    endTime *string,
) *DescribeProbeHistoryRequest {

    return &DescribeProbeHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask/{probeTaskID}/probe/{probeID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ProbeTaskID: probeTaskID,
        ProbeID: probeID,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProbeHistoryRequestWithoutParam() *DescribeProbeHistoryRequest {

    return &DescribeProbeHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask/{probeTaskID}/probe/{probeID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param probeTaskID: 可用性监控task_id, id长度(0,50](Required) */
func (r *DescribeProbeHistoryRequest) SetProbeTaskID(probeTaskID string) {
    r.ProbeTaskID = probeTaskID
}

/* param probeID: 探测源id，  id长度（0,50](Required) */
func (r *DescribeProbeHistoryRequest) SetProbeID(probeID string) {
    r.ProbeID = probeID
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间往前三天，早于3d时，将被重置为3d）(Optional) */
func (r *DescribeProbeHistoryRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，默认为当前时间）(Optional) */
func (r *DescribeProbeHistoryRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProbeHistoryRequest) GetRegionId() string {
    return ""
}

type DescribeProbeHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProbeHistoryResult `json:"result"`
}

type DescribeProbeHistoryResult struct {
    Events []monitor.Event `json:"events"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`
}
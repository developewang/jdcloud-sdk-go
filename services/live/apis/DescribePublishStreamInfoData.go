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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
)

type DescribePublishStreamInfoDataRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    DomainName string `json:"domainName"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 流名称  */
    StreamName string `json:"streamName"`

    /* 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePublishStreamInfoDataRequest(
    domainName string,
    appName string,
    streamName string,
    startTime string,
) *DescribePublishStreamInfoDataRequest {

	return &DescribePublishStreamInfoDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describePublishStreamInfoData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
	}
}

/*
 * param domainName: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 * param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
 (Optional)
 */
func NewDescribePublishStreamInfoDataRequestWithAllParams(
    domainName string,
    appName string,
    streamName string,
    startTime string,
    endTime *string,
) *DescribePublishStreamInfoDataRequest {

    return &DescribePublishStreamInfoDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describePublishStreamInfoData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePublishStreamInfoDataRequestWithoutParam() *DescribePublishStreamInfoDataRequest {

    return &DescribePublishStreamInfoDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describePublishStreamInfoData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 推流域名(Required) */
func (r *DescribePublishStreamInfoDataRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param appName: 应用名称(Required) */
func (r *DescribePublishStreamInfoDataRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param streamName: 流名称(Required) */
func (r *DescribePublishStreamInfoDataRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

/* param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
(Required) */
func (r *DescribePublishStreamInfoDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
(Optional) */
func (r *DescribePublishStreamInfoDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePublishStreamInfoDataRequest) GetRegionId() string {
    return ""
}

type DescribePublishStreamInfoDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePublishStreamInfoDataResult `json:"result"`
}

type DescribePublishStreamInfoDataResult struct {
    DataList []live.PublishStreamInfoResult `json:"dataList"`
}
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

type DescribeLiveTranscodingDurationDataRequest struct {

    core.JDCloudRequest

    /* 码率档次，可以查询指定档次的转码时长，取值：
- video_h264_4k_1
- video_h264_2k_1
- video_h264_fhd_1
- video_h264_hd_1
- video_h264_sd_1
- video_h265_4k_1
- video_h265_2k_1
- video_h265_fhd_1
- video_h265_hd_1
- video_h265_sd_1
 (Optional) */
    Grade *string `json:"grade"`

    /* 查询周期，取值范围：“day,month,year,followTime”，分别表示1天，1月，1年，跟随时间。默认为空，表示day。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional) */
    Period *string `json:"period"`

    /* 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
  */
    StartTime string `json:"startTime"`

    /* 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveTranscodingDurationDataRequest(
    startTime string,
) *DescribeLiveTranscodingDurationDataRequest {

	return &DescribeLiveTranscodingDurationDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeLiveTranscodingDurationData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
	}
}

/*
 * param grade: 码率档次，可以查询指定档次的转码时长，取值：
- video_h264_4k_1
- video_h264_2k_1
- video_h264_fhd_1
- video_h264_hd_1
- video_h264_sd_1
- video_h265_4k_1
- video_h265_2k_1
- video_h265_fhd_1
- video_h265_hd_1
- video_h265_sd_1
 (Optional)
 * param period: 查询周期，取值范围：“day,month,year,followTime”，分别表示1天，1月，1年，跟随时间。默认为空，表示day。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional)
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 * param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
 (Optional)
 */
func NewDescribeLiveTranscodingDurationDataRequestWithAllParams(
    grade *string,
    period *string,
    startTime string,
    endTime *string,
) *DescribeLiveTranscodingDurationDataRequest {

    return &DescribeLiveTranscodingDurationDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodingDurationData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Grade: grade,
        Period: period,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveTranscodingDurationDataRequestWithoutParam() *DescribeLiveTranscodingDurationDataRequest {

    return &DescribeLiveTranscodingDurationDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveTranscodingDurationData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param grade: 码率档次，可以查询指定档次的转码时长，取值：
- video_h264_4k_1
- video_h264_2k_1
- video_h264_fhd_1
- video_h264_hd_1
- video_h264_sd_1
- video_h265_4k_1
- video_h265_2k_1
- video_h265_fhd_1
- video_h265_hd_1
- video_h265_sd_1
(Optional) */
func (r *DescribeLiveTranscodingDurationDataRequest) SetGrade(grade string) {
    r.Grade = &grade
}

/* param period: 查询周期，取值范围：“day,month,year,followTime”，分别表示1天，1月，1年，跟随时间。默认为空，表示day。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
(Optional) */
func (r *DescribeLiveTranscodingDurationDataRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
(Required) */
func (r *DescribeLiveTranscodingDurationDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，为空时默认为当前时间
(Optional) */
func (r *DescribeLiveTranscodingDurationDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveTranscodingDurationDataRequest) GetRegionId() string {
    return ""
}

type DescribeLiveTranscodingDurationDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveTranscodingDurationDataResult `json:"result"`
}

type DescribeLiveTranscodingDurationDataResult struct {
    DataList []live.TranscodeDurationStatisticResult `json:"dataList"`
}
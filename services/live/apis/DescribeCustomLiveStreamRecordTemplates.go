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
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeCustomLiveStreamRecordTemplatesRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1；取值范围[1, 100000] (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 录制模板列表查询过滤条件:
  - name:   template 录制模板自定义名称
  - value:  如果参数为空，则查询全部
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomLiveStreamRecordTemplatesRequest(
) *DescribeCustomLiveStreamRecordTemplatesRequest {

	return &DescribeCustomLiveStreamRecordTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/recordCustoms:template",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码；默认为1；取值范围[1, 100000] (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param filters: 录制模板列表查询过滤条件:
  - name:   template 录制模板自定义名称
  - value:  如果参数为空，则查询全部
 (Optional)
 */
func NewDescribeCustomLiveStreamRecordTemplatesRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeCustomLiveStreamRecordTemplatesRequest {

    return &DescribeCustomLiveStreamRecordTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordCustoms:template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomLiveStreamRecordTemplatesRequestWithoutParam() *DescribeCustomLiveStreamRecordTemplatesRequest {

    return &DescribeCustomLiveStreamRecordTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordCustoms:template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码；默认为1；取值范围[1, 100000](Optional) */
func (r *DescribeCustomLiveStreamRecordTemplatesRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeCustomLiveStreamRecordTemplatesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 录制模板列表查询过滤条件:
  - name:   template 录制模板自定义名称
  - value:  如果参数为空，则查询全部
(Optional) */
func (r *DescribeCustomLiveStreamRecordTemplatesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomLiveStreamRecordTemplatesRequest) GetRegionId() string {
    return ""
}

type DescribeCustomLiveStreamRecordTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomLiveStreamRecordTemplatesResult `json:"result"`
}

type DescribeCustomLiveStreamRecordTemplatesResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    RecordTemplates []live.RecordTemplate `json:"recordTemplates"`
}
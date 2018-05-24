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
    "reflect"
    nc "github.com/jdcloud-api/jdcloud-sdk-go/services/nc/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeSecretsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* name - secret名称，支持模糊搜索
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID 
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: name - secret名称，支持模糊搜索
 (Optional)
 */
func NewDescribeSecretsRequest(
    regionId string,
) *DescribeSecretsRequest {

	return &DescribeSecretsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/secrets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeSecretsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeSecretsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

func (r *DescribeSecretsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

func (r *DescribeSecretsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSecretsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeSecretsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSecretsResult `json:"result"`
}

type DescribeSecretsResult struct {
    Secrets []nc.Secret `json:"secrets"`
    TotalCount int `json:"totalCount"`
}
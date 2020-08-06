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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryPrefetchTaskRequest struct {

    core.JDCloudRequest

    /* url (Optional) */
    Url *string `json:"url"`

    /* 地区[huabei huadong dongbei huazhong huanan xinan xibei gangaotai]中的一个 (Optional) */
    Region *string `json:"region"`

    /* 运营商[ct uni cm]中的一个,分别代表电信 联通 移动 (Optional) */
    Isp *string `json:"isp"`

    /* 查询状态 1:active维护预热中，2:表示purge中暂时停止预热 (Optional) */
    Status *int `json:"status"`

    /* 同url，系统内部url对应id（url和file_id同时存在时以url为准） (Optional) */
    FileId *string `json:"fileId"`

    /* 页码数,最小为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页大小,默认10 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 1:代表控制台下发的预热任务2:代表热度计算下发的预热任务3:代表控制台、热度计算共同下发的任务 (Optional) */
    TaskType *int `json:"taskType"`

    /* 域名 (Optional) */
    Domain *string `json:"domain"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryPrefetchTaskRequest(
) *QueryPrefetchTaskRequest {

	return &QueryPrefetchTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/prefetchTask:query",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param url: url (Optional)
 * param region: 地区[huabei huadong dongbei huazhong huanan xinan xibei gangaotai]中的一个 (Optional)
 * param isp: 运营商[ct uni cm]中的一个,分别代表电信 联通 移动 (Optional)
 * param status: 查询状态 1:active维护预热中，2:表示purge中暂时停止预热 (Optional)
 * param fileId: 同url，系统内部url对应id（url和file_id同时存在时以url为准） (Optional)
 * param pageNumber: 页码数,最小为1 (Optional)
 * param pageSize: 每页大小,默认10 (Optional)
 * param taskType: 1:代表控制台下发的预热任务2:代表热度计算下发的预热任务3:代表控制台、热度计算共同下发的任务 (Optional)
 * param domain: 域名 (Optional)
 */
func NewQueryPrefetchTaskRequestWithAllParams(
    url *string,
    region *string,
    isp *string,
    status *int,
    fileId *string,
    pageNumber *int,
    pageSize *int,
    taskType *int,
    domain *string,
) *QueryPrefetchTaskRequest {

    return &QueryPrefetchTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/prefetchTask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Url: url,
        Region: region,
        Isp: isp,
        Status: status,
        FileId: fileId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        TaskType: taskType,
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryPrefetchTaskRequestWithoutParam() *QueryPrefetchTaskRequest {

    return &QueryPrefetchTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/prefetchTask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param url: url(Optional) */
func (r *QueryPrefetchTaskRequest) SetUrl(url string) {
    r.Url = &url
}

/* param region: 地区[huabei huadong dongbei huazhong huanan xinan xibei gangaotai]中的一个(Optional) */
func (r *QueryPrefetchTaskRequest) SetRegion(region string) {
    r.Region = &region
}

/* param isp: 运营商[ct uni cm]中的一个,分别代表电信 联通 移动(Optional) */
func (r *QueryPrefetchTaskRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param status: 查询状态 1:active维护预热中，2:表示purge中暂时停止预热(Optional) */
func (r *QueryPrefetchTaskRequest) SetStatus(status int) {
    r.Status = &status
}

/* param fileId: 同url，系统内部url对应id（url和file_id同时存在时以url为准）(Optional) */
func (r *QueryPrefetchTaskRequest) SetFileId(fileId string) {
    r.FileId = &fileId
}

/* param pageNumber: 页码数,最小为1(Optional) */
func (r *QueryPrefetchTaskRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页大小,默认10(Optional) */
func (r *QueryPrefetchTaskRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param taskType: 1:代表控制台下发的预热任务2:代表热度计算下发的预热任务3:代表控制台、热度计算共同下发的任务(Optional) */
func (r *QueryPrefetchTaskRequest) SetTaskType(taskType int) {
    r.TaskType = &taskType
}

/* param domain: 域名(Optional) */
func (r *QueryPrefetchTaskRequest) SetDomain(domain string) {
    r.Domain = &domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryPrefetchTaskRequest) GetRegionId() string {
    return ""
}

type QueryPrefetchTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryPrefetchTaskResult `json:"result"`
}

type QueryPrefetchTaskResult struct {
    TotalNumber int `json:"totalNumber"`
    TotalPage int `json:"totalPage"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TaskList []cdn.PrefetchTaskInfo `json:"taskList"`
}
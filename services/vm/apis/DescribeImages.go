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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private、community (Optional) */
    ImageSource *string `json:"imageSource"`

    /* 产品线标识，非必传，不传的时候返回全部产品线镜像 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 是否下线，默认值为false；imageSource为public或者thirdparty时，此参数才有意义，其它情况下此参数无效；指定镜像ID查询时，此参数无效 (Optional) */
    Offline *bool `json:"offline"`

    /* 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu (Optional) */
    Platform *string `json:"platform"`

    /* 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional) */
    Ids []string `json:"ids"`

    /* 镜像支持的系统盘类型，[localDisk,cloudDisk] (Optional) */
    RootDeviceType *string `json:"rootDeviceType"`

    /* 镜像的使用权限[all, specifiedUsers，ownerOnly]，可选参数，仅当imageSource取值private时有效 (Optional) */
    LaunchPermission *string `json:"launchPermission"`

    /* <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a> (Optional) */
    Status *string `json:"status"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImagesRequest(
    regionId string,
) *DescribeImagesRequest {

	return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param imageSource: 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private、community (Optional)
 * param serviceCode: 产品线标识，非必传，不传的时候返回全部产品线镜像 (Optional)
 * param offline: 是否下线，默认值为false；imageSource为public或者thirdparty时，此参数才有意义，其它情况下此参数无效；指定镜像ID查询时，此参数无效 (Optional)
 * param platform: 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu (Optional)
 * param ids: 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional)
 * param rootDeviceType: 镜像支持的系统盘类型，[localDisk,cloudDisk] (Optional)
 * param launchPermission: 镜像的使用权限[all, specifiedUsers，ownerOnly]，可选参数，仅当imageSource取值private时有效 (Optional)
 * param status: <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a> (Optional)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 */
func NewDescribeImagesRequestWithAllParams(
    regionId string,
    imageSource *string,
    serviceCode *string,
    offline *bool,
    platform *string,
    ids []string,
    rootDeviceType *string,
    launchPermission *string,
    status *string,
    pageNumber *int,
    pageSize *int,
) *DescribeImagesRequest {

    return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageSource: imageSource,
        ServiceCode: serviceCode,
        Offline: offline,
        Platform: platform,
        Ids: ids,
        RootDeviceType: rootDeviceType,
        LaunchPermission: launchPermission,
        Status: status,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImagesRequestWithoutParam() *DescribeImagesRequest {

    return &DescribeImagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageSource: 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private、community(Optional) */
func (r *DescribeImagesRequest) SetImageSource(imageSource string) {
    r.ImageSource = &imageSource
}

/* param serviceCode: 产品线标识，非必传，不传的时候返回全部产品线镜像(Optional) */
func (r *DescribeImagesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param offline: 是否下线，默认值为false；imageSource为public或者thirdparty时，此参数才有意义，其它情况下此参数无效；指定镜像ID查询时，此参数无效(Optional) */
func (r *DescribeImagesRequest) SetOffline(offline bool) {
    r.Offline = &offline
}

/* param platform: 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu(Optional) */
func (r *DescribeImagesRequest) SetPlatform(platform string) {
    r.Platform = &platform
}

/* param ids: 镜像ID列表，如果指定了此参数，其它参数可为空(Optional) */
func (r *DescribeImagesRequest) SetIds(ids []string) {
    r.Ids = ids
}

/* param rootDeviceType: 镜像支持的系统盘类型，[localDisk,cloudDisk](Optional) */
func (r *DescribeImagesRequest) SetRootDeviceType(rootDeviceType string) {
    r.RootDeviceType = &rootDeviceType
}

/* param launchPermission: 镜像的使用权限[all, specifiedUsers，ownerOnly]，可选参数，仅当imageSource取值private时有效(Optional) */
func (r *DescribeImagesRequest) SetLaunchPermission(launchPermission string) {
    r.LaunchPermission = &launchPermission
}

/* param status: <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a>(Optional) */
func (r *DescribeImagesRequest) SetStatus(status string) {
    r.Status = &status
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeImagesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeImagesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImagesResult `json:"result"`
}

type DescribeImagesResult struct {
    Images []vm.Image `json:"images"`
    TotalCount int `json:"totalCount"`
}
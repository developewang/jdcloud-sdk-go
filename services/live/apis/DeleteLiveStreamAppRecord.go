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

type DeleteLiveStreamAppRecordRequest struct {

    core.JDCloudRequest

    /* 推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 直播流所属应用名称  */
    AppName string `json:"appName"`

    /* 录制模板自定义名称  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Required)
 * param template: 录制模板自定义名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamAppRecordRequest(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppRecordRequest {

	return &DeleteLiveStreamAppRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/recordApps/{publishDomain}/appNames/{appName}/templates/{template}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        AppName: appName,
        Template: template,
	}
}

/*
 * param publishDomain: 推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Required)
 * param template: 录制模板自定义名称 (Required)
 */
func NewDeleteLiveStreamAppRecordRequestWithAllParams(
    publishDomain string,
    appName string,
    template string,
) *DeleteLiveStreamAppRecordRequest {

    return &DeleteLiveStreamAppRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordApps/{publishDomain}/appNames/{appName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveStreamAppRecordRequestWithoutParam() *DeleteLiveStreamAppRecordRequest {

    return &DeleteLiveStreamAppRecordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordApps/{publishDomain}/appNames/{appName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流加速域名(Required) */
func (r *DeleteLiveStreamAppRecordRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 直播流所属应用名称(Required) */
func (r *DeleteLiveStreamAppRecordRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param template: 录制模板自定义名称(Required) */
func (r *DeleteLiveStreamAppRecordRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamAppRecordRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamAppRecordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamAppRecordResult `json:"result"`
}

type DeleteLiveStreamAppRecordResult struct {
    PublishDomain string `json:"publishDomain"`
}
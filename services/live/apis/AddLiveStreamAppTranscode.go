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

type AddLiveStreamAppTranscodeRequest struct {

    core.JDCloudRequest

    /* 直播的推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 转码模版  */
    Template string `json:"template"`

    /* 直播流所属应用名称  */
    AppName string `json:"appName"`
}

/*
 * param publishDomain: 直播的推流域名 (Required)
 * param template: 转码模版 (Required)
 * param appName: 直播流所属应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddLiveStreamAppTranscodeRequest(
    publishDomain string,
    template string,
    appName string,
) *AddLiveStreamAppTranscodeRequest {

	return &AddLiveStreamAppTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeApps:config",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        Template: template,
        AppName: appName,
	}
}

/*
 * param publishDomain: 直播的推流域名 (Required)
 * param template: 转码模版 (Required)
 * param appName: 直播流所属应用名称 (Required)
 */
func NewAddLiveStreamAppTranscodeRequestWithAllParams(
    publishDomain string,
    template string,
    appName string,
) *AddLiveStreamAppTranscodeRequest {

    return &AddLiveStreamAppTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeApps:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        Template: template,
        AppName: appName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddLiveStreamAppTranscodeRequestWithoutParam() *AddLiveStreamAppTranscodeRequest {

    return &AddLiveStreamAppTranscodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeApps:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 直播的推流域名(Required) */
func (r *AddLiveStreamAppTranscodeRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param template: 转码模版(Required) */
func (r *AddLiveStreamAppTranscodeRequest) SetTemplate(template string) {
    r.Template = template
}

/* param appName: 直播流所属应用名称(Required) */
func (r *AddLiveStreamAppTranscodeRequest) SetAppName(appName string) {
    r.AppName = appName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddLiveStreamAppTranscodeRequest) GetRegionId() string {
    return ""
}

type AddLiveStreamAppTranscodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddLiveStreamAppTranscodeResult `json:"result"`
}

type AddLiveStreamAppTranscodeResult struct {
    PublishDomain string `json:"publishDomain"`
}
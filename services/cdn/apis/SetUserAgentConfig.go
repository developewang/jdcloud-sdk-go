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

type SetUserAgentConfigRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* userAgent类型,取值：block（黑名单）,allow（白名单）,默认为block (Optional) */
    UserAgentType *string `json:"userAgentType"`

    /* UA列表,如果userAgentList为空,则为全部删除 (Optional) */
    UserAgentList []string `json:"userAgentList"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetUserAgentConfigRequest(
    domain string,
) *SetUserAgentConfigRequest {

	return &SetUserAgentConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/userAgentConfig",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param userAgentType: userAgent类型,取值：block（黑名单）,allow（白名单）,默认为block (Optional)
 * param userAgentList: UA列表,如果userAgentList为空,则为全部删除 (Optional)
 */
func NewSetUserAgentConfigRequestWithAllParams(
    domain string,
    userAgentType *string,
    userAgentList []string,
) *SetUserAgentConfigRequest {

    return &SetUserAgentConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/userAgentConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        UserAgentType: userAgentType,
        UserAgentList: userAgentList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetUserAgentConfigRequestWithoutParam() *SetUserAgentConfigRequest {

    return &SetUserAgentConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/userAgentConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetUserAgentConfigRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param userAgentType: userAgent类型,取值：block（黑名单）,allow（白名单）,默认为block(Optional) */
func (r *SetUserAgentConfigRequest) SetUserAgentType(userAgentType string) {
    r.UserAgentType = &userAgentType
}

/* param userAgentList: UA列表,如果userAgentList为空,则为全部删除(Optional) */
func (r *SetUserAgentConfigRequest) SetUserAgentList(userAgentList []string) {
    r.UserAgentList = userAgentList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetUserAgentConfigRequest) GetRegionId() string {
    return ""
}

type SetUserAgentConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetUserAgentConfigResult `json:"result"`
}

type SetUserAgentConfigResult struct {
}
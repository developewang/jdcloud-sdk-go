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

type QueryLiveDomainIpBlackWhiteListRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLiveDomainIpBlackWhiteListRequest(
    domain string,
) *QueryLiveDomainIpBlackWhiteListRequest {

	return &QueryLiveDomainIpBlackWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/ipList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 */
func NewQueryLiveDomainIpBlackWhiteListRequestWithAllParams(
    domain string,
) *QueryLiveDomainIpBlackWhiteListRequest {

    return &QueryLiveDomainIpBlackWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/ipList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryLiveDomainIpBlackWhiteListRequestWithoutParam() *QueryLiveDomainIpBlackWhiteListRequest {

    return &QueryLiveDomainIpBlackWhiteListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/ipList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryLiveDomainIpBlackWhiteListRequest) SetDomain(domain string) {
    r.Domain = domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLiveDomainIpBlackWhiteListRequest) GetRegionId() string {
    return ""
}

type QueryLiveDomainIpBlackWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLiveDomainIpBlackWhiteListResult `json:"result"`
}

type QueryLiveDomainIpBlackWhiteListResult struct {
    Domain string `json:"domain"`
    BlackIps []string `json:"blackIps"`
    WhiteIps []string `json:"whiteIps"`
}
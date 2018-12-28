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

type CreateLiveDomainRequest struct {

    core.JDCloudRequest

    /* 播放域名 (Optional) */
    PlayDomain *string `json:"playDomain"`

    /* 创建推流域名时，必传推流域名 (Optional) */
    PublishDomain *string `json:"publishDomain"`

    /* 回源类型只能是[ips,domain]中的一种 (Optional) */
    SourceType *string `json:"sourceType"`

    /*  (Optional) */
    BackHttpType *string `json:"backHttpType"`

    /* 默认回源host (Optional) */
    DefaultSourceHost *string `json:"defaultSourceHost"`

    /* 站点类型pull(拉流)push(推流) (Optional) */
    SiteType *string `json:"siteType"`

    /* 回源类型，目前只能为rtmp (Optional) */
    BackSourceType *string `json:"backSourceType"`

    /*  (Optional) */
    IpSource []cdn.IpSourceInfo `json:"ipSource"`

    /*  (Optional) */
    DomainSource []cdn.DomainSourceInfo `json:"domainSource"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLiveDomainRequest(
) *CreateLiveDomainRequest {

	return &CreateLiveDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain:batchCreate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param playDomain: 播放域名 (Optional)
 * param publishDomain: 创建推流域名时，必传推流域名 (Optional)
 * param sourceType: 回源类型只能是[ips,domain]中的一种 (Optional)
 * param backHttpType:  (Optional)
 * param defaultSourceHost: 默认回源host (Optional)
 * param siteType: 站点类型pull(拉流)push(推流) (Optional)
 * param backSourceType: 回源类型，目前只能为rtmp (Optional)
 * param ipSource:  (Optional)
 * param domainSource:  (Optional)
 */
func NewCreateLiveDomainRequestWithAllParams(
    playDomain *string,
    publishDomain *string,
    sourceType *string,
    backHttpType *string,
    defaultSourceHost *string,
    siteType *string,
    backSourceType *string,
    ipSource []cdn.IpSourceInfo,
    domainSource []cdn.DomainSourceInfo,
) *CreateLiveDomainRequest {

    return &CreateLiveDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain:batchCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PlayDomain: playDomain,
        PublishDomain: publishDomain,
        SourceType: sourceType,
        BackHttpType: backHttpType,
        DefaultSourceHost: defaultSourceHost,
        SiteType: siteType,
        BackSourceType: backSourceType,
        IpSource: ipSource,
        DomainSource: domainSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLiveDomainRequestWithoutParam() *CreateLiveDomainRequest {

    return &CreateLiveDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain:batchCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param playDomain: 播放域名(Optional) */
func (r *CreateLiveDomainRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = &playDomain
}

/* param publishDomain: 创建推流域名时，必传推流域名(Optional) */
func (r *CreateLiveDomainRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = &publishDomain
}

/* param sourceType: 回源类型只能是[ips,domain]中的一种(Optional) */
func (r *CreateLiveDomainRequest) SetSourceType(sourceType string) {
    r.SourceType = &sourceType
}

/* param backHttpType: (Optional) */
func (r *CreateLiveDomainRequest) SetBackHttpType(backHttpType string) {
    r.BackHttpType = &backHttpType
}

/* param defaultSourceHost: 默认回源host(Optional) */
func (r *CreateLiveDomainRequest) SetDefaultSourceHost(defaultSourceHost string) {
    r.DefaultSourceHost = &defaultSourceHost
}

/* param siteType: 站点类型pull(拉流)push(推流)(Optional) */
func (r *CreateLiveDomainRequest) SetSiteType(siteType string) {
    r.SiteType = &siteType
}

/* param backSourceType: 回源类型，目前只能为rtmp(Optional) */
func (r *CreateLiveDomainRequest) SetBackSourceType(backSourceType string) {
    r.BackSourceType = &backSourceType
}

/* param ipSource: (Optional) */
func (r *CreateLiveDomainRequest) SetIpSource(ipSource []cdn.IpSourceInfo) {
    r.IpSource = ipSource
}

/* param domainSource: (Optional) */
func (r *CreateLiveDomainRequest) SetDomainSource(domainSource []cdn.DomainSourceInfo) {
    r.DomainSource = domainSource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLiveDomainRequest) GetRegionId() string {
    return ""
}

type CreateLiveDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLiveDomainResult `json:"result"`
}

type CreateLiveDomainResult struct {
}
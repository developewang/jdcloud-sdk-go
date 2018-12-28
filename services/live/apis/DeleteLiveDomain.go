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

type DeleteLiveDomainRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`
}

/*
 * param publishDomain: 推流域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveDomainRequest(
    publishDomain string,
) *DeleteLiveDomainRequest {

	return &DeleteLiveDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{publishDomain}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 */
func NewDeleteLiveDomainRequestWithAllParams(
    publishDomain string,
) *DeleteLiveDomainRequest {

    return &DeleteLiveDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{publishDomain}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveDomainRequestWithoutParam() *DeleteLiveDomainRequest {

    return &DeleteLiveDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{publishDomain}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *DeleteLiveDomainRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveDomainRequest) GetRegionId() string {
    return ""
}

type DeleteLiveDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveDomainResult `json:"result"`
}

type DeleteLiveDomainResult struct {
    PublishDomain string `json:"publishDomain"`
}
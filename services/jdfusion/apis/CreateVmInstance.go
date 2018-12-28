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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type CreateVmInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 创建VM  */
    Body *jdfusion.CreateVmReq `json:"body"`
}

/*
 * param regionId: 地域ID (Required)
 * param body: 创建VM (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVmInstanceRequest(
    regionId string,
    body *jdfusion.CreateVmReq,
) *CreateVmInstanceRequest {

	return &CreateVmInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vm_instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Body: body,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param body: 创建VM (Required)
 */
func NewCreateVmInstanceRequestWithAllParams(
    regionId string,
    body *jdfusion.CreateVmReq,
) *CreateVmInstanceRequest {

    return &CreateVmInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Body: body,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVmInstanceRequestWithoutParam() *CreateVmInstanceRequest {

    return &CreateVmInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateVmInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param body: 创建VM(Required) */
func (r *CreateVmInstanceRequest) SetBody(body *jdfusion.CreateVmReq) {
    r.Body = body
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVmInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVmInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVmInstanceResult `json:"result"`
}

type CreateVmInstanceResult struct {
    Task jdfusion.ResourceTFInfo `json:"task"`
}
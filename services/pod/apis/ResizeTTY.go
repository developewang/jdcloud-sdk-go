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

type ResizeTTYRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Pod ID  */
    PodId string `json:"podId"`

    /* container name  */
    ContainerName string `json:"containerName"`

    /* tty row，取值范围：[10, 2000] (Optional) */
    Height *int `json:"height"`

    /* tty column，取值范围：[10, 1000] (Optional) */
    Width *int `json:"width"`

    /* tty column，取值范围：[10, 1000] (Optional) */
    ExecId *string `json:"execId"`
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResizeTTYRequest(
    regionId string,
    podId string,
    containerName string,
) *ResizeTTYRequest {

	return &ResizeTTYRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:resizeTTY",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 * param height: tty row，取值范围：[10, 2000] (Optional)
 * param width: tty column，取值范围：[10, 1000] (Optional)
 * param execId: tty column，取值范围：[10, 1000] (Optional)
 */
func NewResizeTTYRequestWithAllParams(
    regionId string,
    podId string,
    containerName string,
    height *int,
    width *int,
    execId *string,
) *ResizeTTYRequest {

    return &ResizeTTYRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:resizeTTY",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
        Height: height,
        Width: width,
        ExecId: execId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResizeTTYRequestWithoutParam() *ResizeTTYRequest {

    return &ResizeTTYRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:resizeTTY",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ResizeTTYRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param podId: Pod ID(Required) */
func (r *ResizeTTYRequest) SetPodId(podId string) {
    r.PodId = podId
}

/* param containerName: container name(Required) */
func (r *ResizeTTYRequest) SetContainerName(containerName string) {
    r.ContainerName = containerName
}

/* param height: tty row，取值范围：[10, 2000](Optional) */
func (r *ResizeTTYRequest) SetHeight(height int) {
    r.Height = &height
}

/* param width: tty column，取值范围：[10, 1000](Optional) */
func (r *ResizeTTYRequest) SetWidth(width int) {
    r.Width = &width
}

/* param execId: tty column，取值范围：[10, 1000](Optional) */
func (r *ResizeTTYRequest) SetExecId(execId string) {
    r.ExecId = &execId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResizeTTYRequest) GetRegionId() string {
    return r.RegionId
}

type ResizeTTYResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResizeTTYResult `json:"result"`
}

type ResizeTTYResult struct {
}
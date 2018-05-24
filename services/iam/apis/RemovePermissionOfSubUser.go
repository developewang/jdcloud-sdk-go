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
    "reflect"
)

type RemovePermissionOfSubUserRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 权限id  */
    PermissionId int `json:"permissionId"`

    /* 子用户用户名  */
    SubUser string `json:"subUser"`
}

/*
 * param regionId: Region ID 
 * param permissionId: 权限id 
 * param subUser: 子用户用户名 
 */
func NewRemovePermissionOfSubUserRequest(
    regionId string,
    permissionId int,
    subUser string,
) *RemovePermissionOfSubUserRequest {

	return &RemovePermissionOfSubUserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subUser/{subUser}/permissions/{permissionId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PermissionId: permissionId,
        SubUser: subUser,
	}
}

func (r *RemovePermissionOfSubUserRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *RemovePermissionOfSubUserRequest) SetPermissionId(permissionId int) {
    r.PermissionId = permissionId
}

func (r *RemovePermissionOfSubUserRequest) SetSubUser(subUser string) {
    r.SubUser = subUser
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemovePermissionOfSubUserRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type RemovePermissionOfSubUserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemovePermissionOfSubUserResult `json:"result"`
}

type RemovePermissionOfSubUserResult struct {
}
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

type ResetPasswordRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 新密码，必须包含且只支持字母及数字，不少于8字符不超过16字符。  */
    AccountPassword string `json:""`
}

/*
 * param regionId: Region ID 
 * param instanceId: Instance ID 
 * param accountPassword: 新密码，必须包含且只支持字母及数字，不少于8字符不超过16字符。 
 */
func NewResetPasswordRequest(
    regionId string,
    instanceId string,
    accountPassword string,
) *ResetPasswordRequest {

	return &ResetPasswordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:resetPassword",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AccountPassword: accountPassword,
	}
}

func (r *ResetPasswordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ResetPasswordRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ResetPasswordRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResetPasswordRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ResetPasswordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResetPasswordResult `json:"result"`
}

type ResetPasswordResult struct {
}
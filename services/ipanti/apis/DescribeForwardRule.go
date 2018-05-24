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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeForwardRuleRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* 转发规则id  */
    ForwardRuleId string `json:"forwardRuleId"`
}

/*
 * param regionId: Region ID 
 * param instanceId: 实例id 
 * param forwardRuleId: 转发规则id 
 */
func NewDescribeForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleId string,
) *DescribeForwardRuleRequest {

	return &DescribeForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
	}
}

func (r *DescribeForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *DescribeForwardRuleRequest) SetForwardRuleId(forwardRuleId string) {
    r.ForwardRuleId = forwardRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeForwardRuleRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeForwardRuleResult `json:"result"`
}

type DescribeForwardRuleResult struct {
    ForwardRule ipanti.ForwardRule `json:"forwardRule"`
}
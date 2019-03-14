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
    sms "github.com/jdcloud-api/jdcloud-sdk-go/services/sms/models"
)

type SendBatchSmsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 指定模板群发短信请求参数  */
    SendBatchSmsSpec *sms.SendBatchSmsSpec `json:"sendBatchSmsSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param sendBatchSmsSpec: 指定模板群发短信请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendBatchSmsRequest(
    regionId string,
    sendBatchSmsSpec *sms.SendBatchSmsSpec,
) *SendBatchSmsRequest {

	return &SendBatchSmsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/sendBatchSms",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SendBatchSmsSpec: sendBatchSmsSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param sendBatchSmsSpec: 指定模板群发短信请求参数 (Required)
 */
func NewSendBatchSmsRequestWithAllParams(
    regionId string,
    sendBatchSmsSpec *sms.SendBatchSmsSpec,
) *SendBatchSmsRequest {

    return &SendBatchSmsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sendBatchSms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SendBatchSmsSpec: sendBatchSmsSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendBatchSmsRequestWithoutParam() *SendBatchSmsRequest {

    return &SendBatchSmsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sendBatchSms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *SendBatchSmsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param sendBatchSmsSpec: 指定模板群发短信请求参数(Required) */
func (r *SendBatchSmsRequest) SetSendBatchSmsSpec(sendBatchSmsSpec *sms.SendBatchSmsSpec) {
    r.SendBatchSmsSpec = sendBatchSmsSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendBatchSmsRequest) GetRegionId() string {
    return r.RegionId
}

type SendBatchSmsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendBatchSmsResult `json:"result"`
}

type SendBatchSmsResult struct {
    Data []sms.SendBatchSms `json:"data"`
}
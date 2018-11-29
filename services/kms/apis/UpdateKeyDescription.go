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
    kms "github.com/jdcloud-api/jdcloud-sdk-go/services/kms/models"
)

type UpdateKeyDescriptionRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`

    /*   */
    KeyCfg *kms.KeyCfg `json:"keyCfg"`
}

/*
 * param keyId: 密钥ID (Required)
 * param keyCfg:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateKeyDescriptionRequest(
    keyId string,
    keyCfg *kms.KeyCfg,
) *UpdateKeyDescriptionRequest {

	return &UpdateKeyDescriptionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
        KeyCfg: keyCfg,
	}
}

/*
 * param keyId: 密钥ID (Required)
 * param keyCfg:  (Required)
 */
func NewUpdateKeyDescriptionRequestWithAllParams(
    keyId string,
    keyCfg *kms.KeyCfg,
) *UpdateKeyDescriptionRequest {

    return &UpdateKeyDescriptionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
        KeyCfg: keyCfg,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateKeyDescriptionRequestWithoutParam() *UpdateKeyDescriptionRequest {

    return &UpdateKeyDescriptionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *UpdateKeyDescriptionRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

/* param keyCfg: (Required) */
func (r *UpdateKeyDescriptionRequest) SetKeyCfg(keyCfg *kms.KeyCfg) {
    r.KeyCfg = keyCfg
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateKeyDescriptionRequest) GetRegionId() string {
    return ""
}

type UpdateKeyDescriptionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateKeyDescriptionResult `json:"result"`
}

type UpdateKeyDescriptionResult struct {
}
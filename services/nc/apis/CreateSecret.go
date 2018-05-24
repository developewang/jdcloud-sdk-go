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
    nc "github.com/jdcloud-api/jdcloud-sdk-go/services/nc/models"
)

type CreateSecretRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 机密数据名称，不能重复  */
    Name string `json:"name"`

    /* 私密数据的类型，目前仅支持如下类型：docker-registry：用来和docker registry认证的类型  */
    SecretType string `json:"secretType"`

    /* 机密的数据 (Optional) */
    Data *nc.DockerRegistryData `json:"data"`
}

/*
 * param regionId: Region ID 
 * param name: 机密数据名称，不能重复 
 * param secretType: 私密数据的类型，目前仅支持如下类型：docker-registry：用来和docker registry认证的类型 
 * param data: 机密的数据 (Optional)
 */
func NewCreateSecretRequest(
    regionId string,
    name string,
    secretType string,
) *CreateSecretRequest {

	return &CreateSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/secrets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        SecretType: secretType,
	}
}

func (r *CreateSecretRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *CreateSecretRequest) SetName(name string) {
    r.Name = name
}

func (r *CreateSecretRequest) SetSecretType(secretType string) {
    r.SecretType = secretType
}

func (r *CreateSecretRequest) SetData(data *nc.DockerRegistryData) {
    r.Data = data
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSecretRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type CreateSecretResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSecretResult `json:"result"`
}

type CreateSecretResult struct {
    SecretName string `json:"secretName"`
}
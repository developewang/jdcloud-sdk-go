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

type ImportThingModelRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 产品Key  */
    ProductKey string `json:"productKey"`

    /* 物模型JSON  */
    ThingModel interface{} `json:"thingModel"`
}

/*
 * param regionId: 地域ID (Required)
 * param productKey: 产品Key (Required)
 * param thingModel: 物模型JSON (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImportThingModelRequest(
    regionId string,
    productKey string,
    thingModel interface{},
) *ImportThingModelRequest {

	return &ImportThingModelRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/products/{productKey}/abilities:importThingModel",
			Method:  "PUT",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        ProductKey: productKey,
        ThingModel: thingModel,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param productKey: 产品Key (Required)
 * param thingModel: 物模型JSON (Required)
 */
func NewImportThingModelRequestWithAllParams(
    regionId string,
    productKey string,
    thingModel interface{},
) *ImportThingModelRequest {

    return &ImportThingModelRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/products/{productKey}/abilities:importThingModel",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        ProductKey: productKey,
        ThingModel: thingModel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImportThingModelRequestWithoutParam() *ImportThingModelRequest {

    return &ImportThingModelRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/products/{productKey}/abilities:importThingModel",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ImportThingModelRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productKey: 产品Key(Required) */
func (r *ImportThingModelRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param thingModel: 物模型JSON(Required) */
func (r *ImportThingModelRequest) SetThingModel(thingModel interface{}) {
    r.ThingModel = thingModel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImportThingModelRequest) GetRegionId() string {
    return r.RegionId
}

type ImportThingModelResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImportThingModelResult `json:"result"`
}

type ImportThingModelResult struct {
}
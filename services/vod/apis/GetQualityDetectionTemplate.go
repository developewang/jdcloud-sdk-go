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

type GetQualityDetectionTemplateRequest struct {

    core.JDCloudRequest

    /* 模板ID  */
    TemplateId int `json:"templateId"`
}

/*
 * param templateId: 模板ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetQualityDetectionTemplateRequest(
    templateId int,
) *GetQualityDetectionTemplateRequest {

	return &GetQualityDetectionTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/qualityDetectionTemplates/{templateId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TemplateId: templateId,
	}
}

/*
 * param templateId: 模板ID (Required)
 */
func NewGetQualityDetectionTemplateRequestWithAllParams(
    templateId int,
) *GetQualityDetectionTemplateRequest {

    return &GetQualityDetectionTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionTemplates/{templateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TemplateId: templateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetQualityDetectionTemplateRequestWithoutParam() *GetQualityDetectionTemplateRequest {

    return &GetQualityDetectionTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionTemplates/{templateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateId: 模板ID(Required) */
func (r *GetQualityDetectionTemplateRequest) SetTemplateId(templateId int) {
    r.TemplateId = templateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetQualityDetectionTemplateRequest) GetRegionId() string {
    return ""
}

type GetQualityDetectionTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetQualityDetectionTemplateResult `json:"result"`
}

type GetQualityDetectionTemplateResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    TemplateType string `json:"templateType"`
    Detections []string `json:"detections"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}
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

type QueryDefaultHttpHeaderKeyRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDefaultHttpHeaderKeyRequest(
) *QueryDefaultHttpHeaderKeyRequest {

	return &QueryDefaultHttpHeaderKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/defaultHttpHeaderKey",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewQueryDefaultHttpHeaderKeyRequestWithAllParams(
) *QueryDefaultHttpHeaderKeyRequest {

    return &QueryDefaultHttpHeaderKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/defaultHttpHeaderKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDefaultHttpHeaderKeyRequestWithoutParam() *QueryDefaultHttpHeaderKeyRequest {

    return &QueryDefaultHttpHeaderKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/defaultHttpHeaderKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDefaultHttpHeaderKeyRequest) GetRegionId() string {
    return ""
}

type QueryDefaultHttpHeaderKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDefaultHttpHeaderKeyResult `json:"result"`
}

type QueryDefaultHttpHeaderKeyResult struct {
    ReqDefaultHttpHeaderKey []string `json:"reqDefaultHttpHeaderKey"`
    RespDefaultHttpHeaderKey []string `json:"respDefaultHttpHeaderKey"`
}
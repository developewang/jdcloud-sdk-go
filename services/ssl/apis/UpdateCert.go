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

type UpdateCertRequest struct {

    core.JDCloudRequest

    /* 证书ID  */
    CertId string `json:"certId"`

    /* 私钥  */
    KeyFile string `json:"keyFile"`

    /* 证书  */
    CertFile string `json:"certFile"`
}

/*
 * param certId: 证书ID (Required)
 * param keyFile: 私钥 (Required)
 * param certFile: 证书 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCertRequest(
    certId string,
    keyFile string,
    certFile string,
) *UpdateCertRequest {

	return &UpdateCertRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslCert:updateCert",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CertId: certId,
        KeyFile: keyFile,
        CertFile: certFile,
	}
}

/*
 * param certId: 证书ID (Required)
 * param keyFile: 私钥 (Required)
 * param certFile: 证书 (Required)
 */
func NewUpdateCertRequestWithAllParams(
    certId string,
    keyFile string,
    certFile string,
) *UpdateCertRequest {

    return &UpdateCertRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:updateCert",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CertId: certId,
        KeyFile: keyFile,
        CertFile: certFile,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCertRequestWithoutParam() *UpdateCertRequest {

    return &UpdateCertRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:updateCert",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param certId: 证书ID(Required) */
func (r *UpdateCertRequest) SetCertId(certId string) {
    r.CertId = certId
}

/* param keyFile: 私钥(Required) */
func (r *UpdateCertRequest) SetKeyFile(keyFile string) {
    r.KeyFile = keyFile
}

/* param certFile: 证书(Required) */
func (r *UpdateCertRequest) SetCertFile(certFile string) {
    r.CertFile = certFile
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCertRequest) GetRegionId() string {
    return ""
}

type UpdateCertResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCertResult `json:"result"`
}

type UpdateCertResult struct {
    CertId string `json:"certId"`
    Digest string `json:"digest"`
}
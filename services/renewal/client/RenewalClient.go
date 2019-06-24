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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    renewal "github.com/jdcloud-api/jdcloud-sdk-go/services/renewal/apis"
    "encoding/json"
    "errors"
)

type RenewalClient struct {
    core.JDCloudClient
}

func NewRenewalClient(credential *core.Credential) *RenewalClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("renewal.jdcloud-api.com")

    return &RenewalClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "renewal",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *RenewalClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *RenewalClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 开通、取消实例自动续费 */
func (c *RenewalClient) SetRenewal(request *renewal.SetRenewalRequest) (*renewal.SetRenewalResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &renewal.SetRenewalResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 实例续费 */
func (c *RenewalClient) RenewInstance(request *renewal.RenewInstanceRequest) (*renewal.RenewInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &renewal.RenewInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可续费实例 */
func (c *RenewalClient) QueryInstance(request *renewal.QueryInstanceRequest) (*renewal.QueryInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &renewal.QueryInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}


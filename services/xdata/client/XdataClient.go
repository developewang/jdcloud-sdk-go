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
    xdata "github.com/jdcloud-api/jdcloud-sdk-go/services/xdata/apis"
    "encoding/json"
    "errors"
)

type XdataClient struct {
    core.JDCloudClient
}

func NewXdataClient(credential *core.Credential) *XdataClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("xdata.jdcloud-api.com")

    return &XdataClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "xdata",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *XdataClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *XdataClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 删除数据表 */
func (c *XdataClient) DeleteTable(request *xdata.DeleteTableRequest) (*xdata.DeleteTableResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.DeleteTableResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询指定数据库下所有数据表 */
func (c *XdataClient) ListTableInfo(request *xdata.ListTableInfoRequest) (*xdata.ListTableInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListTableInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询实例列表 */
func (c *XdataClient) ListDatabaseInfo(request *xdata.ListDatabaseInfoRequest) (*xdata.ListDatabaseInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListDatabaseInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除数据库 */
func (c *XdataClient) DeleteDatabase(request *xdata.DeleteDatabaseRequest) (*xdata.DeleteDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.DeleteDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建数据库 */
func (c *XdataClient) CreateDatabase(request *xdata.CreateDatabaseRequest) (*xdata.CreateDatabaseResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CreateDatabaseResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询实例列表 */
func (c *XdataClient) ListInstanceInfo(request *xdata.ListInstanceInfoRequest) (*xdata.ListInstanceInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.ListInstanceInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询数据库详情 */
func (c *XdataClient) GetDatabaseInfo(request *xdata.GetDatabaseInfoRequest) (*xdata.GetDatabaseInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetDatabaseInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询数据表信息 */
func (c *XdataClient) GetTableInfo(request *xdata.GetTableInfoRequest) (*xdata.GetTableInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.GetTableInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建数据表 */
func (c *XdataClient) CreateTable(request *xdata.CreateTableRequest) (*xdata.CreateTableResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &xdata.CreateTableResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}


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

package models


type RedisType struct {

    /* type名称  */
    Name string `json:"name"`

    /* 当前type的key个数  */
    KeyNumber int64 `json:"keyNumber"`

    /* 当前type的key的内存大小，单位：字节，redis2.8无此项  */
    KeySize int64 `json:"keySize"`

    /* type占比  */
    Ratio float32 `json:"ratio"`
}

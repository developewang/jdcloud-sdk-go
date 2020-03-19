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


type AudioScanResultDetail struct {

    /* 句子开始的时间，单位是秒。 (Optional) */
    StartTime int `json:"startTime"`

    /* 该句语言的检测结果的分类，取值参见audioScenes与label参数说明。 (Optional) */
    EndTime string `json:"endTime"`

    /* 语音转换成文本的结果。 (Optional) */
    Text string `json:"text"`

    /* 命中该风险的上下文信息。具体结构描述见hintWordsInfo (Optional) */
    HintWordsInfos []HintWordsInfo `json:"hintWordsInfos"`
}

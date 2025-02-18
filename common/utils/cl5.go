/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package utils

import (
	"errors"
	"fmt"
	"github.com/polarismesh/polaris-server/common/model"
	"strconv"
	"strings"
)

// cl5集群的ctx的key
type Cl5ServerCluster struct{}

// cl5.server的协议ctx
type Cl5ServerProtocol struct{}

// Sid结构体，序列化转为sid字符串
func MarshalSid(sid *model.Sid) string {
	return fmt.Sprintf("%d:%d", sid.ModID, sid.CmdID)
}

// mod cmd转为sid
func MarshalModCmd(modID uint32, cmdID uint32) string {
	return fmt.Sprintf("%d:%d", modID, cmdID)
}

// 把sid字符串反序列化为结构体Sid
func UnmarshalSid(sidStr string) (*model.Sid, error) {
	items := strings.Split(sidStr, ":")
	if len(items) != 2 {
		return nil, errors.New("invalid format sid string")
	}

	modID, err := strconv.ParseUint(items[0], 10, 32)
	if err != nil {
		return nil, err
	}
	cmdID, err := strconv.ParseUint(items[1], 10, 32)
	if err != nil {
		return nil, err
	}

	out := &model.Sid{
		ModID: uint32(modID),
		CmdID: uint32(cmdID),
	}
	return out, nil
}

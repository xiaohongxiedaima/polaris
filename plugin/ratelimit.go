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

package plugin

import (
	"github.com/polarismesh/polaris-server/common/log"
	"os"
	"sync"
)

type RatelimitType int

const (
	// 基于ip的流控
	IPRatelimit RatelimitType = iota + 1

	// 基于接口级限流
	APIRatelimit

	// 基于service的流控
	ServiceRatelimit

	// 基于Instance的限流
	InstanceRatelimit
)

var RatelimitStr = map[RatelimitType]string{
	IPRatelimit:       "ip-limit",
	APIRatelimit:      "api-limit",
	ServiceRatelimit:  "service-limit",
	InstanceRatelimit: "instance-limit",
}

var (
	ratelimitOnce = sync.Once{}
)

/**
 * @brief Ratelimit插件接口
 */
type Ratelimit interface {
	Plugin

	// 是否允许访问, true: 允许, false: 不允许 TODO
	// 参数rateType即限流类型，id为限流的key
	// 如果rateType为RatelimitIP则id为ip, rateType为RatelimitService则id为ip_namespace_service或ip_serviceId
	Allow(typ RatelimitType, key string) bool
}

/**
 * @brief 获取RateLimit插件
 */
func GetRatelimit() Ratelimit {
	c := &config.RateLimit

	plugin, exist := pluginSet[c.Name]
	if !exist {
		return nil
	}

	ratelimitOnce.Do(func() {
		if err := plugin.Initialize(c); err != nil {
			log.Errorf("plugin init err: %s", err.Error())
			os.Exit(-1)
		}
	})

	return plugin.(Ratelimit)
}

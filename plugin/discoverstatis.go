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
	"time"
)

var (
	discoverStatisOnce = &sync.Once{}
)

/**
 * @brief 服务发现统计插件接口
 */
type DiscoverStatis interface {
	Plugin

	AddDiscoverCall(service, namespace string, time time.Time) error
}

/**
 * @brief 获取服务发现统计插件
 */
func GetDiscoverStatis() DiscoverStatis {
	c := &config.DiscoverStatis

	plugin, exist := pluginSet[c.Name]
	if !exist {
		return nil
	}

	discoverStatisOnce.Do(func() {
		if err := plugin.Initialize(c); err != nil {
			log.Errorf("plugin init err: %s", err.Error())
			os.Exit(-1)
		}
	})

	return plugin.(DiscoverStatis)
}

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

package boltdbStore

import "github.com/polarismesh/polaris-server/common/model"

type platformStore struct {
	handler BoltHandler
}

// 新增平台信息
func (p *platformStore) CreatePlatform(platform *model.Platform) error {
	//TODO
	return nil
}

// 更新平台信息
func (p *platformStore) UpdatePlatform(platform *model.Platform) error {
	//TODO
	return nil
}

// 删除平台信息
func (p *platformStore) DeletePlatform(id string) error {
	//TODO
	return nil
}

// 查询平台信息
func (p *platformStore) GetPlatformById(id string) (*model.Platform, error) {
	//TODO
	return nil, nil
}

// 根据过滤条件查询平台信息
func (p *platformStore) GetPlatforms(
	query map[string]string, offset uint32, limit uint32) (uint32, []*model.Platform, error) {
	//TODO
	return 0, nil, nil
}

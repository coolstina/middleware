// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redisclient

import (
	"testing"
	"time"

	redisc "github.com/coolstina/connecter/redis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestRedisClientSuite(t *testing.T) {
	suite.Run(t, &RedisClientSuite{})
}

type RedisClientSuite struct {
	suite.Suite
	err         error
	redisclient *redis.Client
}

func (suite *RedisClientSuite) BeforeTest(suiteName, testName string) {
	suite.redisclient, suite.err = NewRedisClientOr(redisc.NewDefaultSimpleConfig("localhost:6379", "", 3))
	assert.NoError(suite.T(), suite.err)
	assert.NotNil(suite.T(), suite.redisclient)
}

func (suite *RedisClientSuite) Test_SetGet() {
	// Set key value.
	set := suite.redisclient.Set("username", "helloshaohua", time.Second*60)
	assert.NotNil(suite.T(), set)
	actual, err := set.Result()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "OK", actual)

	// Get key value.
	actual, err = suite.redisclient.Get("username").Result()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "helloshaohua", actual)
}

func (suite *RedisClientSuite) Test_Get_NotExists() {
	// Get key value.
	result, err := suite.redisclient.Exists("not_exists_key").Result()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), int64(0), result)
}

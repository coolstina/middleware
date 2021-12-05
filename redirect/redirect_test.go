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

package redirect

import (
	"net/http"
	"testing"

	redisc "github.com/coolstina/connecter/redis"
	"github.com/coolstina/middleware/connector/redisclient"
	"github.com/go-redis/redis"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestRedirectSuite(t *testing.T) {
	suite.Run(t, &RedirectSuite{})
}

type RedirectSuite struct {
	suite.Suite
	err             error
	client          *redis.Client
	eventsWithFile  EmitTriggerEvent
	eventsWithRedis EmitTriggerEvent
}

func (suite *RedirectSuite) BeforeTest(suiteName, testName string) {
	suite.client, suite.err = redisclient.NewRedisClientOr(redisc.NewDefaultSimpleConfig("localhost:6379", "", 3))
	assert.NoError(suite.T(), suite.err)
	assert.NotNil(suite.T(), suite.client)

	suite.eventsWithFile = EmitTriggerEvent{
		"/usr/local/data/fss/.keep.upgrading": &TriggerEvent{
			StatusCode:            http.StatusServiceUnavailable,
			RedirectURI:           "/fss/upgrading",
			RedirectHeaderMessage: "FSS application upgrading...",
		},
	}

	suite.eventsWithFile = EmitTriggerEvent{
		".keep.upgrading": &TriggerEvent{
			StatusCode:            http.StatusServiceUnavailable,
			RedirectURI:           "/fss/upgrading",
			RedirectHeaderMessage: "FSS application upgrading...",
		},
	}
}

func (suite *RedirectSuite) Test_NewRedirect_WithMonitor_File() {
	redirect := NewRedirect(
		suite.eventsWithFile,
		WatcherOfFile,
	)
	assert.NotNil(suite.T(), redirect)
}

func (suite *RedirectSuite) Test_NewRedirect_WithMonitor_Redis() {
	redirect := NewRedirect(
		suite.eventsWithRedis,
		WatcherOfRedis,
		WithWatcherOfRedisClient(suite.client),
	)
	assert.NotNil(suite.T(), redirect)
}

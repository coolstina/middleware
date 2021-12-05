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
	"github.com/coolstina/fsfire"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// DefaultRedirectMessageHeader Default HTTP response header constant definition for redirect message.
const (
	DefaultRedirectMessageHeader = "X-Redirect-Message"
)

type Redirecter interface {
	Handler() gin.HandlerFunc
}

type Redirect struct {
	emits   Emits
	watcher Watcher
	options *options
	client  *redis.Client
}

func NewRedirect(emits Emits, watcher Watcher, ops ...Option) *Redirect {
	options := &options{}
	for _, o := range ops {
		o.apply(options)
	}

	if options.header == "" {
		options.header = DefaultRedirectMessageHeader
	}

	if watcher == WatcherOfRedis && options.client == nil {
		panic("use redis watcher please configuration the redis client connection instance")
	}

	return &Redirect{emits: emits, watcher: watcher, options: options, client: options.client}
}

func (redirect *Redirect) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		if redirect.emits.URIIsRedirect(uri) {
			switch redirect.watcher {
			case WatcherOfFile:
				for monitor, event := range redirect.emits {
					if !fsfire.IsNotExists(monitor) {
						ctx.Writer.Header().Set(redirect.options.header, event.RedirectHeaderMessage)
						ctx.Redirect(event.StatusCode, event.RedirectURI)
						ctx.Abort()
					}
				}
			case WatcherOfRedis:
				for monitor, event := range redirect.emits {
					exists, _ := redirect.client.Exists(monitor).Result()
					if exists == int64(1) {
						ctx.Writer.Header().Set(redirect.options.header, event.RedirectHeaderMessage)
						ctx.Redirect(event.StatusCode, event.RedirectURI)
						ctx.Abort()
					}
				}
			}
		}
		ctx.Next()
	}
}

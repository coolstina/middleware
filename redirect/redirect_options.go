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

import "github.com/go-redis/redis"

// Option design modern impl.
type Option interface {
	apply(*options)
}

type optionFunc func(ops *options)

func (of optionFunc) apply(ops *options) {
	of(ops)
}

type options struct {
	client *redis.Client
	header string
}

// WithWatcherOfRedisClient Specify redisclient config when watcher is redisclient.
func WithWatcherOfRedisClient(client *redis.Client) Option {
	return optionFunc(func(ops *options) {
		ops.client = client
	})
}

// WithRedirectMessageHeader Specify redirect message HTTP response header.
func WithRedirectMessageHeader(header string) Option {
	return optionFunc(func(ops *options) {
		ops.header = header
	})
}

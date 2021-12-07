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
	"github.com/stretchr/testify/assert"
)

func (suite *RedirectSuite) Test_Options() {
	options := &options{}

	ops := []Option{
		WithSetIgnoreURI("/users/helloshaohua/info", true),
		WithSetIgnoreURI("/users/helloshaohua/status", true),
	}

	for _, o := range ops {
		o.apply(options)
	}
	assert.Len(suite.T(), options.ignore, 2)
}

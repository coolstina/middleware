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

import "github.com/stretchr/testify/assert"

func (suite *RedirectSuite) Test_Ignore_URIIsIgnore_Want_False() {
	var ignore Ignore
	actual := ignore.URIIsIgnore("hello world")
	assert.Equal(suite.T(), false, actual)
}

func (suite *RedirectSuite) Test_Ignore_URIIsIgnore_Want_True() {
	var ignore = Ignore{"/users/helloshaohua/info": true}
	actual := ignore.URIIsIgnore("/users/helloshaohua/info")
	assert.Equal(suite.T(), true, actual)
}

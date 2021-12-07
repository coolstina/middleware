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

// Ignore will skip the redirect detection without triggering redirection events.
type Ignore map[string]bool

// URIIsIgnore Checks whether the URI is set to ignore.
func (ignore Ignore) URIIsIgnore(uri string) bool {
	access.Lock()
	defer access.Unlock()

	for k, v := range ignore {
		if k == uri && v {
			return true
		}
	}
	return false
}

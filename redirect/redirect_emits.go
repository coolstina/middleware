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
	"sync"
)

// Emits is a mapping configuration used to trigger events.
type Emits map[string]*TriggerEvent

var access sync.Mutex

// TriggerEvent Get the trigger event for key.
func (emit Emits) TriggerEvent(key string) *TriggerEvent {
	access.Lock()
	defer access.Unlock()

	if val, ok := emit[key]; ok {
		return val
	}
	return nil
}

// SetTriggerEvent If the triggering event already exists, it is overwritten.
func (emit Emits) SetTriggerEvent(key string, event *TriggerEvent) Emits {
	access.Lock()
	defer access.Unlock()

	emit[key] = event
	return emit
}

func (emit Emits) redirects() map[string]string {
	var uris = make(map[string]string)
	for _, event := range emit {
		uris[event.RedirectURI] = event.RedirectURI
	}
	return uris
}

func (emit Emits) URIIsRedirect(uri string) bool {
	access.Lock()
	defer access.Unlock()

	redirects := emit.redirects()
	if _, exists := redirects[uri]; exists {
		return false
	}
	return true
}

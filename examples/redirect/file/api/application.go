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

package api

import (
	"net/http"

	"github.com/coolstina/fsfire"
	"github.com/gin-gonic/gin"
)

type API struct{}

func (api *API) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "welcome to fss application"})
}

func (api *API) Upgrade(ctx *gin.Context) {
	file := "/usr/local/data/fss/.keep.upgrading"
	if fsfire.IsNotExists(file) {
		err := fsfire.Touch(file)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "the application is about to start upgrading"},
			)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "the application is about to start upgrading"})
}

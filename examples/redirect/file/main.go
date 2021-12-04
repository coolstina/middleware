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

package main

import (
	"log"

	"github.com/coolstina/middleware/examples/redirect/file/api"
	"github.com/coolstina/middleware/examples/redirect/file/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	redirect := middleware.GetRedirectOr()
	handler := &api.API{}
	engine := gin.Default()
	engine.Use(redirect.Handler())
	engine.GET("/welcome", handler.Welcome)
	engine.GET("/upgrading", handler.Upgrade)
	log.Fatal(engine.Run(":9090"))
}

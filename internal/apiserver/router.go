// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	"github.com/ngsin/iam-learning/internal/apiserver/store/mysql"
	"github.com/ngsin/iam-learning/internal/pkg/middleware/auth"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	basicStrategy := newBasicAuth().(auth.BasicStrategy)
	g.POST("/login", basicStrategy.AuthFunc())
	store.Client()
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	fmt.Printf("\nstoreIns: %v\n", storeIns)
	//TODO: define controller
	return g
}

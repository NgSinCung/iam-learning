// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/ngsin/iam-learning/internal/apiserver/controller/v1/policy"
	"github.com/ngsin/iam-learning/internal/apiserver/controller/v1/secret"
	"github.com/ngsin/iam-learning/internal/apiserver/controller/v1/user"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	"github.com/ngsin/iam-learning/internal/pkg/code"
	"github.com/ngsin/iam-learning/internal/pkg/middleware"
	"github.com/ngsin/iam-learning/internal/pkg/middleware/auth"

	// custom gin validators.
	_ "github.com/ngsin/iam-learning/internal/pkg/validator"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	// Refresh time can be longer than token timeout
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	// no route match run auto auth
	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})

	// get store instance
	storeIns := store.Client()

	// router
	v1 := g.Group("/v1")
	{
		v1.Use(auto.AuthFunc())

		// user RESTful resource
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.Use(auto.AuthFunc(), middleware.Validation())
			userv1.POST("", userController.Create)
			//userv1.DELETE("", userController.DeleteCollection) // admin api
			//userv1.DELETE(":name", userController.Delete)      // admin api
			//userv1.PUT(":name/change-password", userController.ChangePassword)
			//userv1.PUT(":name", userController.Update)
			userv1.GET("", userController.List)
			userv1.GET(":name", userController.Get) // admin api

		}

		// policy RESTful resource
		policyv1 := v1.Group("/policies")
		{
			policyController := policy.NewPolicyController(storeIns)
			policyv1.GET(":name", policyController.Get)
		}

		// secret RESTful resource
		secretv1 := v1.Group("/secrets")
		{
			secretController := secret.NewSecretController(storeIns)
			secretv1.GET("", secretController.List)
			secretv1.GET(":name", secretController.Get)

		}
	}

	return g
}

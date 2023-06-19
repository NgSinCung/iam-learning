// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import "github.com/gin-gonic/gin"

// GenericAPIServer contains a custom gin.Engine and some middlewares
type GenericAPIServer struct {
	*gin.Engine
	middlewares       []string
	SecureServingInfo *SecureServingInfo
}

func (s GenericAPIServer) Setup() {

}

func (s GenericAPIServer) InstallMiddlewares() {

}

func (s GenericAPIServer) InstallAPIs() {

}

func initGenericAPIServer(s *GenericAPIServer) {
	// do some setup
	// s.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// GenericAPIServer contains a custom gin.Engine and some middlewares
type GenericAPIServer struct {
	*gin.Engine
	middlewares                  []string
	SecureServingInfo            *SecureServingInfo
	InsecureServingInfo          *InsecureServingInfo
	insecureServer, secureServer *http.Server
}

func (s *GenericAPIServer) Setup() {

}

func (s *GenericAPIServer) InstallMiddlewares() {

}

func (s *GenericAPIServer) InstallAPIs() {

}

func initGenericAPIServer(s *GenericAPIServer) {
	// do some setup
	// s.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	//TODO: run gin server
	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.InsecureServingInfo.Address,
		Handler: s,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,

	}

	// For scalability, use custom HTTP configuration mode here
	s.secureServer = &http.Server{
		Addr:    s.SecureServingInfo.Address(),
		Handler: s,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	eg.Go(func() error {
		// TODO: fmt to log
		fmt.Printf("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf(err.Error())

			return err
		}

		fmt.Printf("Server on %s stopped", s.InsecureServingInfo.Address)

		return nil
	})

	//TODO: health check

	if err := eg.Wait(); err != nil {
		// TODO: fmt to log
		fmt.Printf(err.Error())
	}

	return nil
}

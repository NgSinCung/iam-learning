// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github/ngsin/iam-learning/internal/apiserver/config"
	genericapiserver "github/ngsin/iam-learning/internal/pkg/apiserver"
)

type Server struct {
	genericAPIServer *genericapiserver.GenericAPIServer
}

type PreparedServer struct {
	*Server
}

func (s *PreparedServer) Run() error {
	return s.genericAPIServer.Run()
}

func (s *Server) PrepareRun() *PreparedServer {
	initRouter(s.genericAPIServer.Engine)

	return &PreparedServer{s}

}

func createServer(genericConfig *config.Config) (*Server, error) {
	genericServerConfig, err := buildGenericServerConfig(genericConfig)
	if err != nil {
		return nil, err
	}
	genericServer := genericServerConfig.Complete().NewServer()
	server := &Server{
		genericAPIServer: genericServer,
		// TODO: add other server here. ex. grpc server
	}
	return server, nil
}

func buildGenericServerConfig(genericConfig *config.Config) (genericServerConfig *genericapiserver.Config, lastErr error) {
	// generate api server default required config
	genericServerConfig = genericapiserver.NewConfig()

	if lastErr = genericConfig.GenericServerRunOptions.ApplyTo(genericServerConfig); lastErr != nil {
		return
	}

	if lastErr = genericConfig.SecureServingOptions.ApplyTo(genericServerConfig); lastErr != nil {
		return
	}

	if lastErr = genericConfig.InsecureServingOptions.ApplyTo(genericServerConfig); lastErr != nil {
		return
	}
	return
}

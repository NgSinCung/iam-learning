// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"GoPracticalDevelopmentDemo/internel/apiserver/config"
	genericapiserver "GoPracticalDevelopmentDemo/internel/pkg/server"
)

type apiServer struct {
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func (s *preparedAPIServer) Run() error {
	return s.genericAPIServer.Run()
}

func (s *apiServer) PrepareRun() *preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	return &preparedAPIServer{s}

}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericServer := genericConfig.Complete().New()
	server := &apiServer{
		genericAPIServer: genericServer,
	}
	return server, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return nil, lastErr
	}
	return
}

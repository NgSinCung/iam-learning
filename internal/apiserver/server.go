// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/ngsin/iam-learning/internal/apiserver/config"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	"github.com/ngsin/iam-learning/internal/apiserver/store/mysql"
	"github.com/ngsin/iam-learning/internal/pkg/api/rest"
	genericoptions "github.com/ngsin/iam-learning/internal/pkg/options"
)

type Server struct {
	genericRESTServer *rest.GenericAPIServer
	extraConfig       *ExtraConfig
}

type PreparedServer struct {
	*Server
}

func (s *PreparedServer) Run() error {
	return s.genericRESTServer.Run()
}

func (s *Server) PrepareRun() *PreparedServer {
	storeIns, _ := mysql.GetMySQLFactoryOr(s.extraConfig.mysqlOptions)
	// storeIns, _ := etcd.GetEtcdFactoryOr(c.etcdOptions, nil)
	store.SetClient(storeIns)

	initRouter(s.genericRESTServer.Engine)

	return &PreparedServer{s}

}

func createServer(genericConfig *config.Config) (*Server, error) {
	genericRESTServerConfig, err := buildGenericRESTServerConfig(genericConfig)
	if err != nil {
		return nil, err
	}

	extraConfig, _ := buildExtraConfig(genericConfig)
	if err != nil {
		return nil, err
	}

	genericServer := genericRESTServerConfig.Complete().NewServer()

	extraConfig.complete()

	server := &Server{
		genericRESTServer: genericServer,
		extraConfig:       extraConfig,
		// TODO: add other server here. ex. grpc server
	}
	return server, nil
}

func buildGenericRESTServerConfig(genericConfig *config.Config) (genericRESTServerConfig *rest.Config, lastErr error) {
	// generate api server default required config
	genericRESTServerConfig = rest.NewConfig()

	if lastErr = genericConfig.GenericServerRunOptions.ApplyTo(genericRESTServerConfig); lastErr != nil {
		return
	}

	if lastErr = genericConfig.SecureServingOptions.ApplyTo(genericRESTServerConfig); lastErr != nil {
		return
	}

	if lastErr = genericConfig.InsecureServingOptions.ApplyTo(genericRESTServerConfig); lastErr != nil {
		return
	}
	return
}

// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	ServerCert   genericoptions.GeneratableKeyCert
	mysqlOptions *genericoptions.MySQLOptions
	// etcdOptions      *genericoptions.EtcdOptions
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		mysqlOptions: cfg.MySQLOptions,
	}, nil
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}
}

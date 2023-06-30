// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/marmotedu/iam/pkg/log"
	"github.com/ngsin/iam-learning/internal/apiserver/config"
	"github.com/ngsin/iam-learning/internal/apiserver/store"
	"github.com/ngsin/iam-learning/internal/apiserver/store/mysql"
	"github.com/ngsin/iam-learning/internal/pkg/api/rest"
	genericoptions "github.com/ngsin/iam-learning/internal/pkg/options"
	"github.com/ngsin/iam-learning/pkg/shutdown"
	"github.com/ngsin/iam-learning/pkg/shutdown/shutdownmanagers/posixsignal"
)

type Server struct {
	gs                *shutdown.GracefulShutdown
	genericRESTServer *rest.GenericAPIServer
	extraConfig       *ExtraConfig
}

type PreparedServer struct {
	*Server
}

func (s *PreparedServer) Run() error {

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericRESTServer.Run()
}

func (s *Server) PrepareRun() *PreparedServer {

	storeIns, err := mysql.GetMySQLFactoryOr(s.extraConfig.mysqlOptions)
	if err != nil {
		log.Fatalf("get mysql factory failed: %s", err.Error())
		return nil
	}
	// storeIns, _ := etcd.GetEtcdFactoryOr(c.etcdOptions, nil)
	store.SetClient(storeIns)

	initRouter(s.genericRESTServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
			log.Infof("shutdown callback: %s", "close mysql store")
		}

		// TODO: add grpc server graceful shutdown
		//s.gRPCAPIServer.Close()
		//s.genericAPIServer.Close()

		return nil
	}))

	return &PreparedServer{s}

}

func createServer(genericConfig *config.Config) (*Server, error) {
	// graceful shutdown
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

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
		gs:                gs,
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

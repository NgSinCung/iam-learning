// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import "github.com/gin-gonic/gin"

type Config struct {
	Mode          string
	SecureServing *SecureServingInfo
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// NewConfig return a default config
func NewConfig() *Config {
	return &Config{
		Mode: gin.ReleaseMode,
	}
}

type CompletedConfig struct {
	*Config
}

func (c CompletedConfig) New() *GenericAPIServer {
	gin.SetMode(c.Mode)
	s := &GenericAPIServer{
		Engine:            gin.New(),
		SecureServingInfo: c.SecureServing,
	}
	initGenericAPIServer(s)
	return s
}

// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

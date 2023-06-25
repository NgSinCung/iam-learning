// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/gin-gonic/gin"
)

// Config is a structure used to configure a GenericAPIServer.
// Its members are sorted roughly in order of importance for composers.
// Its members refer to the services included on the server.
type Config struct {
	Mode            string
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Mode: gin.ReleaseMode,
	}
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}

// NewServer return a GenericAPIServer
func (c CompletedConfig) NewServer() *GenericAPIServer {
	gin.SetMode(c.Mode)
	s := &GenericAPIServer{
		Engine:              gin.New(),
		SecureServingInfo:   c.SecureServing,
		InsecureServingInfo: c.InsecureServing,
	}
	initGenericAPIServer(s)
	return s
}

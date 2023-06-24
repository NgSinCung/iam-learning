// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github/ngsin/iam-learning/internal/apiserver/config"
)

// Run runs the specified APIServer.
func Run(genericConfig *config.Config) error {
	server, err := createServer(genericConfig)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run()
}

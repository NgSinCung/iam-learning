// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"fmt"
	"github/ngsin/iam-learning/internel/apiserver/config"
)

func Run(cfg *config.Config) error {
	fmt.Println("run apiserver")
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run()
}

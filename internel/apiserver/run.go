// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"GoPracticalDevelopmentDemo/internel/apiserver/config"
	"fmt"
)

func Run(cfg *config.Config) error {
	fmt.Println("run apiserver")
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run()
}

// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import "github/NgSingCung/iam-learning/internel/pkg/server"

type ServerRunOptions struct {
	Mode string `json:"mode"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()
	return &ServerRunOptions{
		Mode: defaults.Mode,
	}
}

func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	return nil
}

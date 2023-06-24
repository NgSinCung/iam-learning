// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ngsin/iam-learning/internal/pkg/apiserver"
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	Mode string `json:"mode"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := apiserver.NewConfig()
	return &ServerRunOptions{
		Mode: defaults.Mode,
	}
}

func (s *ServerRunOptions) ApplyTo(c *apiserver.Config) error {
	c.Mode = s.Mode
	return nil
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs.StringVar(&s.Mode, "server.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")
}

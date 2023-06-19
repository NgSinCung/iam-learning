// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github/NgSingCung/iam-learning/internel/apiserver/config"
	"github/NgSingCung/iam-learning/internel/apiserver/options"
	"github/NgSingCung/iam-learning/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithDescription("IAM API Server"),
		app.WithRunFunc(run(opts)),
		app.WithOptions(opts),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		//todo log something
		//todo init config
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}

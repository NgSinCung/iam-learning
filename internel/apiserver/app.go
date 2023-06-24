// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github/ngsin/iam-learning/internel/apiserver/config"
	"github/ngsin/iam-learning/internel/apiserver/options"
	"github/ngsin/iam-learning/pkg/app"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)
	return application
}

// run return a func which is used as the RunFunc of app.App.
func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		//TODO: log something
		//TODO: init config
		genericConfig, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(genericConfig)
	}
}

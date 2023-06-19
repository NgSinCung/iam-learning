// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	genericoptions "github/NgSingCung/iam-learning/internel/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	SecureServing           *genericoptions.SecureServingOptions
}

func NewOptions() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
	}
}

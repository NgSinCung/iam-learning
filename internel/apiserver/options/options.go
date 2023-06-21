// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	genericoptions "github/ngsin/iam-learning/internel/pkg/options"
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

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	return fss
}

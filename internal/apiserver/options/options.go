// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	genericoptions "github.com/ngsin/iam-learning/internal/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	InsecureServingOptions  *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServingOptions    *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
}

func NewOptions() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		SecureServingOptions:    genericoptions.NewSecureServingOptions(),
		InsecureServingOptions:  genericoptions.NewInsecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
	}
}

// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/marmotedu/iam/pkg/log"
	genericoptions "github.com/ngsin/iam-learning/internal/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	InsecureServingOptions  *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServingOptions    *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
}

func NewOptions() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		SecureServingOptions:    genericoptions.NewSecureServingOptions(),
		InsecureServingOptions:  genericoptions.NewInsecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		Log:                     log.NewOptions(),
	}
}

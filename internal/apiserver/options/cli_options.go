// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// implement the CliOptions interface

package options

import cliflag "github.com/marmotedu/component-base/pkg/cli/flag"

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServingOptions.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServingOptions.AddFlags(fss.FlagSet("secure serving"))
	return fss
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error
	return errs
}

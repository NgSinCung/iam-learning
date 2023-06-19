// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github/ngsin/iam-learning/internel/apiserver"
	"math/rand"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	apiserver.NewApp("iam-apiserver").Run()
}

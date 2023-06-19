// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/spf13/cobra"
)

type App struct {
	basename    string
	name        string
	description string
	runFunc     RunFunc
	options     CliOptions
	args        cobra.PositionalArgs
}

type RunFunc func(basename string) error

func WithRunFunc(runFunc RunFunc) func(*App) {
	return func(a *App) {
		a.runFunc = runFunc
	}
}

func WithDescription(description string) func(*App) {
	return func(a *App) {
		a.description = description
	}
}

type Option func(*App)

func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

// NewApp is used to create a new application.
func NewApp(name string, basename string, options ...func(*App)) *App {
	app := &App{
		name:     name,
		basename: basename,
	}

	for _, option := range options {
		option(app)
	}

	app.buildCommand()

	return app
}

// Run is used to launch the application.
func (a *App) Run() error {
	return a.runFunc(a.basename)
}

func (a *App) buildCommand() {

}

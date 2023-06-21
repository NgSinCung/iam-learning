// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"github.com/fatih/color"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/component-base/pkg/cli/globalflag"
	"github.com/marmotedu/component-base/pkg/term"
	"github.com/marmotedu/component-base/pkg/version/verflag"
	"github.com/spf13/cobra"
	"os"
)

// App is the main structure of a cli application.
// It is recommended that an app be created with the app.NewApp() function.
type App struct {
	basename    string // basename is the name of the binary.
	name        string // name is the name of the application.
	description string
	runFunc     RunFunc              // runFunc is the function to be executed when the application is launched.
	options     CliOptions           // options is the options of the application.
	args        cobra.PositionalArgs // args is a func type to be used for validating positional arguments.
	commands    []*Command           // commands is the subcommands that are part of this application.
	cmd         *cobra.Command       // cmd is the cobra command which use runFunc as RunE.
	noVersion   bool                 // noVersion is used to determine whether to display version information.
	noConfig    bool                 // noConfig is used to determine whether to read configuration information.
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
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func (a *App) buildCommand() {
	// create the root command
	cmd := cobra.Command{
		Use:   FormatBaseName(a.basename),
		Short: a.name,
		Long:  a.description,
		// stop printing usage when the command errors
		SilenceUsage: true,
		// stop printing error when the command errors
		SilenceErrors: true,
		Args:          a.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	// add global flags to the command and normalize them
	cliflag.InitFlags(cmd.Flags())

	// add subcommands to the command
	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
		cmd.SetHelpCommand(helpCommand(FormatBaseName(a.basename)))
	}

	// add run function to the command
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	// get flags from options and add them to the command
	var namedFlagSets cliflag.NamedFlagSets
	if a.options != nil {
		// get flags from options
		namedFlagSets = a.options.Flags()
		// add flags to the command FlagSet
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	// add version and config flags to the namedFlagSets
	if !a.noVersion {
		verflag.AddFlags(namedFlagSets.FlagSet("global"))
	}

	// add config flag to the namedFlagSets
	if !a.noConfig {
		addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))
	}

	// add help flag to the namedFlagSets
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())

	// add namedFlagSets to the command FlagSet
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))

	// set UsageFunc and HelpFunc for the command
	addCmdTemplate(&cmd, namedFlagSets)

	// set the command to the app
	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	//todo: init job before run runFunc
	fmt.Println("init job before run runFunc")

	// if pass --version flag, print version information and exit
	if !a.noVersion {
		verflag.PrintAndExitIfRequested()
	}

	// run application
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}
	return nil
}

// addCmdTemplate is used to set UsageFunc and HelpFunc for the command.
func addCmdTemplate(cmd *cobra.Command, namedFlagSets cliflag.NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		// print usage
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		// print flags
		cliflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}

package app

import (
	"fmt"
	"github.com/marmotedu/component-base/pkg/util/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const configFlagName = "config"

var cfgFile string

func init() {
	// add config flag which is the path of configuration file
	pflag.StringVarP(&cfgFile, configFlagName, "c", cfgFile, "Read configuration from specified `FILE`, "+
		"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

// addConfigFlag adds flags for a specific server to the specified FlagSet
// object.
func addConfigFlag(basename string, fs *pflag.FlagSet) {

	// add config flag to the specified FlagSet
	fs.AddFlag(pflag.Lookup(configFlagName))

	// auto load config from env variables
	viper.AutomaticEnv()
	// set prefix for env variables and replace - with _ in given prefix string
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))
	// replace char in key name
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// the given function will be called before any command's Execute method is called.
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			// add current directory as config path
			viper.AddConfigPath(".")

			// add home directory as config path
			if names := strings.Split(basename, "-"); len(names) > 1 {
				viper.AddConfigPath(filepath.Join(homedir.HomeDir(), "."+names[0]))
				viper.AddConfigPath(filepath.Join("/etc", names[0]))
			}

			viper.SetConfigName(basename)
		}

		// read config file
		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}
	})
}

package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "flixinit",
		Short: "Flixinit is a simple CLI tool to make your application a great tenant for cloud environments",
		Long: `Flixinit is a simple CLI tool to make your application a great tenant for cloud environments.
Complete documentation is available at https://github.com/saeedafshari8/flixinit`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flixinit.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Saeed Afshari", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "", "Apache 2.0", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.SetDefault("author", "Saeed Afshari <saeed.afshari8@gmail.com>")
	viper.SetDefault("license", "Apache 2.0")

	rootCmd.AddCommand(cmdJava)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		util.LogAndExit(err, util.EnvironmentError)

		// Search config in home directory with name ".flixinit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".flixinit")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

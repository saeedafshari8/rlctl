package cmd

import (
	"fmt"
	"github.com/rocketlaunchercloud/rlctl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "rlctl",
		Short: "Rlctl is a simple CLI tool to make your application a great tenant for cloud environments",
		Long: `RlCtl is a simple CLI tool to make your application a great tenant for cloud environments.
Complete documentation is available at https://github.com/rocketlaunchercloud/rlctl`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Run rlctl -h for help.")
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	initFlags()

	rootCmd.AddCommand(SpringCommand)
	rootCmd.AddCommand(cmdGitLab)
}

func initFlags() {
	rootCmd.PersistentFlags().StringP("author", "a", "Saeed Afshari", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "", "Apache 2.0", "name of license for the project")
}

func initConfig() {
	util.InitConfig()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file:%v\n", viper.ConfigFileUsed())
	}

	viper.SetDefault("author", "Saeed Afshari <saeed.afshari8@gmail.com>")
	viper.SetDefault("license", "Apache 2.0")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

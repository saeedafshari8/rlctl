package util

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	GitlabToken = "gitlab-token"
)

func GetGitlabToken() string {
	return viper.GetString(GitlabToken)
}

func SetGitlabToken(token string) {
	viper.Set(GitlabToken, token)
	err := viper.WriteConfig()
	if err != nil {
		LogAndExit(err, FileNotFound)
	}
}

func InitConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	LogAndExit(err, EnvironmentError)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".flixinit")

	configName := ".flixinit.yaml"
	configPath := path.Join(home, configName)
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		viper.Set(GitlabToken, "")
		var data []byte
		err = ioutil.WriteFile(configPath, data, os.ModePerm)
		LogAndExit(err, EnvironmentError)
	}
	if err != nil {
		log.Printf("Unable to created config %s\n", configPath)
		LogAndExit(err, EnvironmentError)
	}
}

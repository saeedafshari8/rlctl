package util

import (
	"fmt"
	"github.com/spf13/cobra"
)

func ValidateRequired(value, key string) {
	if value == "" {
		LogMessageAndExit(fmt.Sprintf("%s is mandatory!\n", key))
	}
}

func GetValue(cmd *cobra.Command, key string) string {
	s, err := cmd.Flags().GetString(key)
	LogAndExit(err, ArgMissing)
	return s
}

func GetValueBool(cmd *cobra.Command, key string) bool {
	b, err := cmd.Flags().GetBool(key)
	LogAndExit(err, ArgMissing)
	return b
}

func GetValueInt32(cmd *cobra.Command, key string) int32 {
	value, err := cmd.Flags().GetInt32(key)
	LogAndExit(err, ArgMissing)
	return value
}

func GetValues(cmd *cobra.Command, key string) []string {
	b, err := cmd.Flags().GetStringArray(key)
	LogAndExit(err, ArgMissing)
	return b
}

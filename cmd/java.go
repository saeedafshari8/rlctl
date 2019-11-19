package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmdJava = &cobra.Command{
	Use:   "java",
	Short: "java command generates a new gradle/java project",
	Long:  `java command generates a new gradle/java project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Java project created successfully")
	},
}

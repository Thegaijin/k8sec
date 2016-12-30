package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootOpts = struct {
	debug      bool
	kubeconfig string
	namespace  string
}{}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "k8sec",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if rootOpts.debug {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Println(err)
		}
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().BoolVar(&rootOpts.debug, "debug", false, "Debug mode")
	RootCmd.PersistentFlags().StringVar(&rootOpts.kubeconfig, "kubeconfig", "", "Path of kubeconfig file")
	RootCmd.PersistentFlags().StringVar(&rootOpts.namespace, "namespace", "", "Namespace scope")
}

func initConfig() {
}

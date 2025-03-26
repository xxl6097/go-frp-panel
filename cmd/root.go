package cmd

import (
	"fmt"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/spf13/cobra"
	"github.com/xxl6097/go-frp-panel/pkg"
	"os"
)

var (
	showVersion bool
	showSecret  bool
	run         func() error
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version of frps-panel")
	rootCmd.PersistentFlags().BoolVarP(&showSecret, "secret", "s", false, "display frps-panel config")
}

var rootCmd = &cobra.Command{
	Use:   "frps-panel",
	Short: "frps-panel is the server plugin of frp to support multiple users.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println("Frps Version:\t", version.Full())
			pkg.Version()
			return nil
		}
		if showSecret {
			//config.ShowConfig()
		}
		if run != nil {
			return run()
		}
		return nil
	},
}

func Execute(f func() error) {
	run = f
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

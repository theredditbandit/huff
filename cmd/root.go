/*
Copyright © 2024 theredditbandit
*/
package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "huff",
	Short: "A huffman encoder/decoder",
	Run: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if debug && err == nil {
			log.SetLevel(log.DebugLevel)
			log.Info("huff is running in debug mode . . .")
		} else {
			log.SetLevel(log.InfoLevel)
			log.Info("huff is running in info mode . . .")
		}
		log.SetFormatter(&log.TextFormatter{
            ForceColors: true,
        })
        // log.SetReportCaller(true)
		log.SetOutput(os.Stdout)
        log.WithField("Args", args).Debug("arguments")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode")
}

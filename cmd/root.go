package cmd

import (
	"fmt"
	"os"

	"huff/pkg"

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
		log.SetOutput(os.Stdout)
		log.WithField("Args", args).Debug("arguments")

		if len(args) != 0 {
			fname := args[0]
			data, err := os.ReadFile(fname)
			if err != nil {
				log.WithField("File Name:", fname).Fatal("Not a valid file")
			} else {
				log.WithField("File ", fname).Info("File read successfully")
			}
			stats := pkg.GetStats(data)
			logStats(stats)
            
		}
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

func logStats(stats map[string]int) {
	log.Debug("File stats: ")
	for k, v := range stats {
		log.Debug(fmt.Sprintf("%s: %d", k, v))
	}
}

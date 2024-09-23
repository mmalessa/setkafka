package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"setkafka/pkg/app"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "setkafka",
		Short: "Set Kafka",
	}
	verbose  bool
	quiet    bool
	cfgFiles []string = []string{"./setkafka.yaml", "/etc/setkafka/setkafka.yaml"}
	cfgFile  string   = ""
)

func Execute() error {
	cobra.OnInitialize(initConfig, initLogs)
	return rootCmd.Execute()
}

func init() {
	cfgFilesString := strings.Join(cfgFiles, ", ")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", fmt.Sprintf("config file (default: %s)", cfgFilesString))
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose (log Trace level)")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Be quiet (log Error level)")
}

func initConfig() {
	var err error
	absCfgPath, err := getConfigFileName(cfgFile)
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}

	if !quiet {
		fmt.Printf("Config file used: %s\n", absCfgPath)
	}
	if err := app.InitConfig(absCfgPath); err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
}

func getConfigFileName(cfgFile string) (string, error) {
	if cfgFile != "" {
		absFn, err := filepath.Abs(cfgFile)
		if err != nil {
			return "", err
		}
		return absFn, nil
	}

	for _, filename := range cfgFiles {
		absFn, err := filepath.Abs(filename)
		if err != nil {
			continue
		}
		return absFn, nil
	}
	return "", fmt.Errorf("config file not found")
}

func initLogs() {
	llvl := ""
	if verbose {
		llvl = "trace"
	} else if quiet {
		llvl = "error"
	}
	if err := app.InitLogs(llvl); err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
}

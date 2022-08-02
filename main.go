package main

import (
	"github.com/git-zjx/git-local-stats/stats"
	"github.com/spf13/cobra"
	"log"
)

var CmdScan = &cobra.Command{
	Use:   "scan",
	Short: "Scan an folder",
	Long:  "Scan that folder and its subdirectories for repositories to scan, Example: git-stats scan /path/to/folder",
	Run:   scanRun,
}

var CmdStats = &cobra.Command{
	Use:   "stats",
	Short: "Generate a CLI stats graph for the passed email",
	Long:  "Generate a CLI stats graph representing the last 6 months of activity for the passed email, Example: git-stats stats 977904037@qq.com",
	Run:   statsRun,
}

func scanRun(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		return
	}
	for _, arg := range args {
		stats.Scan(arg)
	}
}

func statsRun(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		return
	}
	email := args[0]
	stats.Print(email)
}

var rootCmd = &cobra.Command{
	Use:     "git-stats",
	Short:   "git-stats: An tool for Git local stats.",
	Long:    `git-stats: An tool for Git local stats.`,
	Version: "1.0",
}

func init() {
	rootCmd.AddCommand(CmdScan)
	rootCmd.AddCommand(CmdStats)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

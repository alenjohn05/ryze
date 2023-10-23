package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/alenjohn05/ryze/cmd/apps"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var all bool = false

var rootCmd = &cobra.Command{
	Use:   "ryze [email]",
	Short: "ryze helps you find apps where target mail is registered. [v0.1]",
	Args:  cobra.ExactArgs(1),
	Long:  "ryze helps you find apps where target mail is registered. [v0.1]",
	Run: func(cmd *cobra.Command, args []string) {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(args[0]) {
			color.Red("Invalid email address!")
			os.Exit(0)
		}
		start := time.Now()
		fmt.Println("Where are you", color.YellowString(args[0]), "? \U0001f9d0")

		var wg sync.WaitGroup
		wg.Add(6)
		go apps.Instagram(&wg, args[0], all)
		go apps.Twitter(&wg, args[0], all)
		go apps.Spotify(&wg, args[0], all)
		go apps.Adobe(&wg, args[0], all)
		go apps.Vsco(&wg, args[0], all)
		go apps.Discord(&wg, args[0], all)
		wg.Wait()

		elapsed := time.Since(start)
		fmt.Println("The scan took", color.GreenString(fmt.Sprintf("%.2f", elapsed.Seconds())), "seconds.")
		os.Exit(0)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVar(&all, "all", false, "Show all results")
}

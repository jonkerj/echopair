package cmd

import (
	"log"
	"os"

	"github.com/jonkerj/echopair/pkg/client"
	"github.com/spf13/cobra"
)

var (
	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "simple tcp client",
		Run:   hammer,
	}

	remoteAddr string
	interval   string
)

func init() {
	flags := clientCmd.PersistentFlags()
	flags.StringVarP(&interval, "interval", "i", "5s", "interval between echos")
	flags.StringVarP(&remoteAddr, "remote", "r", "", "host to echo with")

	rootCmd.AddCommand(clientCmd)

}

func hammer(cobra *cobra.Command, args []string) {
	s := client.New(remoteAddr, log.New(os.Stdout, "", log.LstdFlags), interval)
	s.EchoMany()
}

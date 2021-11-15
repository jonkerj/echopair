package cmd

import (
	"log"
	"os"

	"github.com/jonkerj/echopair/pkg/server"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "a simple tcp echo server",
		Run:   serve,
	}

	tcpPort int
)

func init() {
	flags := serveCmd.Flags()
	flags.IntVarP(&tcpPort, "port", "p", 25455, "TCP port to use")

	rootCmd.AddCommand(serveCmd)
}

func serve(cobra *cobra.Command, args []string) {
	s := server.New(tcpPort, log.New(os.Stdout, "", log.LstdFlags))
	s.Serve()
}

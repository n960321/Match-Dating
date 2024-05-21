package cmd

import (
	"Match-Dating/internal/config"
	"Match-Dating/internal/handler"
	"Match-Dating/internal/server"
	"Match-Dating/pkg/logger"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run:   RunServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func RunServer(cmd *cobra.Command, args []string) {
	logger.SetLogger(local)
	config := config.GetConfig(configFile)

	handler := handler.NewHandler()
	svr := server.NewServer(config.Http, handler)

	svr.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	svr.Shutdown(ctx)

	log.Info().Msg("shutting down")
	os.Exit(0)

}

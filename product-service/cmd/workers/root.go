package workers

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hailsayan/sophocles/pkg/logger"
	"github.com/hailsayan/sophocles/product-service/internal/config"
	"github.com/hailsayan/sophocles/product-service/internal/log"
	"github.com/hailsayan/sophocles/product-service/internal/provider"
	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewLogrusLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(cfg, ctx)
			},
			PreRun: func(cmd *cobra.Command, args []string) {
				go runKafkaWorker(cfg, ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}

package workers

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hailsayan/sophocles/order-service/internal/config"
	"github.com/hailsayan/sophocles/order-service/internal/log"
	"github.com/hailsayan/sophocles/order-service/internal/provider"
	"github.com/hailsayan/sophocles/pkg/logger"
	"github.com/spf13/cobra"
)

func Start() {
	cfg := config.InitConfig()
	log.SetLogger(logger.NewLogrusLogger(cfg.Logger.Level))
	provider.BootstrapGlobal(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(cfg, ctx, &wg)
			},
			PreRun: func(cmd *cobra.Command, args []string) {
				go func() {
					wg.Wait()
					go runKafkaWorker(cfg, ctx)
					go runAMQPWorker(cfg, ctx)
				}()
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}

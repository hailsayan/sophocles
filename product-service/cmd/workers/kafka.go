package workers

import (
	"context"

	"github.com/hailsayan/sophocles/product-service/internal/config"
	"github.com/hailsayan/sophocles/product-service/internal/server"
)

func runKafkaWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewKafkaServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

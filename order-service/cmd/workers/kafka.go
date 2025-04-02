package workers

import (
	"context"

	"github.com/hailsayan/sophocles/order-service/internal/config"
	"github.com/hailsayan/sophocles/order-service/internal/server"
)

func runKafkaWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewKafkaServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

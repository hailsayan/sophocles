package workers

import (
	"context"

	"github.com/hailsayan/sophocles/mail-service/internal/config"
	"github.com/hailsayan/sophocles/mail-service/internal/server"
)

func runAMQPWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewAMQPServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

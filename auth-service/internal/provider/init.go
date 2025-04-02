package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/hailsayan/sophocles/auth-service/internal/config"
	"github.com/hailsayan/sophocles/auth-service/internal/repository"
	"github.com/hailsayan/sophocles/pkg/database"
	"github.com/hailsayan/sophocles/pkg/utils/encryptutils"
	"github.com/hailsayan/sophocles/pkg/utils/jwtutils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	db           *sqlx.DB
	rdb          *redis.ClusterClient
	rabbitmq     *amqp.Connection
	bcryptHasher encryptutils.Hasher
	jwtUtil      jwtutils.JwtUtil
	dataStore    repository.DataStore
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rdb = database.NewRedisCluster((*database.RedisClusterOptions)(cfg.Redis))
	rabbitmq = database.NewAMQP((*database.AmqpOptions)(cfg.Amqp))
	bcryptHasher = encryptutils.NewBcryptHasher(cfg.App.BCryptCost)
	jwtUtil = jwtutils.NewJwtUtil()
	dataStore = repository.NewDataStore(db)
}

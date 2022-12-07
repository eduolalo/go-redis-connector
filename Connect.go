/*
	En este documento se escribe la función necesaria para establecer una
   conexión con la base de datos REDIS
*/
package redis

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var mainClient *redis.Client
var mu sync.Mutex

// Connect proporciona la funcionalidad necesaria para establecer una conexión
// en chinga y que administre el pool y esas cosas
func Connect() *redis.Client {

	mu.Lock()
	defer mu.Unlock()
	if mainClient != nil {
		return mainClient
	}
	connString, ok := os.LookupEnv("RDS_STRING")
	if !ok {
		log.Panic("Could not find RDS_STRING environment variable")
	}
	if connString == "" {
		log.Panic(`RDS_STRING environment variable cannot be empty`)
	}
	opt, err := redis.ParseURL(connString)
	if err != nil {
		log.Panicf("Error on parsing connection string %+v", err)
	}

	opt.PoolSize = poolSize()
	client := redis.NewClient(opt)

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Panicf("Error on ping database, %+v", err)
	}

	mainClient = client
	setConnectionName()
	return mainClient
}

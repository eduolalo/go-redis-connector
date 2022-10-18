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

	"github.com/mediocregopher/radix/v4"
)

var connString = os.Getenv("RDS_STRING")

var mainClient *radix.Client
var mu sync.Mutex

// Connect proporciona la funcionalidad necesaria para establecer una conexión
// en chinga y que administre el pool y esas cosas
func Connect() radix.Client {

	mu.Lock()
	defer mu.Unlock()
	if mainClient != nil {
		return *mainClient
	}

	config := radix.PoolConfig{
		Size: poolSize(),
	}
	log.Println("*** Creando nuevo pool de Redis ***")
	client, err := config.New(context.Background(), "tcp", connString)
	if err != nil {

		log.Println("****** No se pudo establecer conexión con redis ******")
		log.Println(err.Error())
		log.Println("****** No se pudo establecer conexión con redis ******")
		return nil
	}
	mainClient = &client
	go setConnectionName()
	return *mainClient
}

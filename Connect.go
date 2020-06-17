/*
	En este documento se escribe la función necesaria para establecer una
   conexión con la base de datos REDIS
*/
package redis

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

var connString = os.Getenv("RDS_STRING")

var mainPool *radix.Pool

// Connect proporciona la funcionalidad necesaria para establecer una conexión
// en chinga y que administre el pool y esas cosas
func Connect() (pool *radix.Pool) {

	if mainPool == nil {

		log.Println("*** Creando nuevo pool de Redis ***")
		pool, err := radix.NewPool("tcp", connString, poolSize())
		if err != nil {

			log.Println("****** No se pudo establecer conexión con redis ******")
			log.Println(err.Error())
			log.Println("****** No se pudo establecer conexión con redis ******")
			return nil
		}
		mainPool = pool
		go setConnectionName()
	}
	return mainPool
}

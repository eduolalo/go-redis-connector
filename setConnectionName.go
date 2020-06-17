package redis

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

// setConnectionName es para signar nombre a la conexión de Redis
func setConnectionName() {

	name := os.Getenv("RDS_NAME")
	var resp string
	err := mainPool.Do(radix.Cmd(&resp, "CLIENT", "SETNAME", name))

	if err != nil {

		log.Println("****** Error redis.setName ******")
		log.Println(err)
		log.Println("------ Error redis.setName ------")
		return
	}

	log.Println("****** Redis client name: ", name, " ******")
	log.Printf(resp)
}

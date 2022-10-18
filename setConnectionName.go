package redis

import (
	"context"
	"log"
	"os"

	"github.com/mediocregopher/radix/v4"
)

// setConnectionName es para signar nombre a la conexión de Redis
func setConnectionName() {

	if mainClient == nil {
		return
	}
	client := *mainClient
	name, ok := os.LookupEnv("RDS_NAME")
	if !ok {
		if hostName, err := os.Hostname(); err == nil {
			name = hostName
		}
	}
	var resp string
	err := client.Do(context.Background(), radix.Cmd(&resp, "CLIENT", "SETNAME", name))

	if err != nil {

		log.Println("****** Error redis.setName ******")
		log.Println(err)
		log.Println("------ Error redis.setName ------")
		return
	}

	log.Println("****** Redis client name: ", name, " ******")
}

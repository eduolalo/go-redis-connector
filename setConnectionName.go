package redis

import (
	"context"
	"log"
	"os"
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
	_, err := client.Do(context.Background(), "CLIENT", "SETNAME", name).Result()
	if err != nil {
		log.Printf("Error on set connection name: %+v", err)
		return
	}
}

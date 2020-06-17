package redis

import (
	"os"
	"strconv"
)

// poolSize regresa el número de conexiónes configuradas para el sistema
func poolSize() (connections int) {

	conn := os.Getenv("RDS_POOL_SIZE")
	connections, err := strconv.Atoi(conn)
	if err != nil {
		connections = 3
	}
	return
}

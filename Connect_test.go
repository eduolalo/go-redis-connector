package redis

import (
	"context"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {

	t.Run("Redis_Lib", func(t *testing.T) {

		connString, ok := os.LookupEnv("RDS_STRING")
		if !ok {
			t.Fatal(`No se encontró la variable "RDS_STRING" necesaria para la conexión`)
			return
		}
		if connString == "" {
			t.Fatal(`La variable "RDS_STRING" no puede estar vacía`)
			return
		}

		t.Run("Echo", func(t *testing.T) {
			pool := Connect()
			ctx := context.Background()

			var val string = "jaló esta merga"
			if str, err := pool.Echo(ctx, val).Result(); err != nil {

				t.Fatal(err)
			} else if str != val {

				t.Errorf("el valor ejecutado %s es diferente al recibido %s", val, str)
			}
		})

		t.Run("Connection Name", func(t *testing.T) {
			pool := Connect()
			ctx := context.Background()

			name, ok := os.LookupEnv("RDS_NAME")
			if !ok {
				if hostName, err := os.Hostname(); err == nil {
					name = hostName
				}
			}
			conn := pool.Conn(ctx)
			if nm, err := conn.ClientGetName(ctx).Result(); err != nil {
				t.Fatal(err)
			} else if nm != name {

				t.Errorf("Name spected: %s, name received %s", name, nm)
			}
		})
	})
}

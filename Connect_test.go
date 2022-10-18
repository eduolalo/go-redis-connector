package redis

import (
	"context"
	"os"
	"testing"

	"github.com/mediocregopher/radix/v4"
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

		t.Run("Connect", func(t *testing.T) {
			pool := Connect()
			defer pool.Close()

			var str, val string
			val = "jaló esta merga"
			ctx := context.Background()
			if err := pool.Do(ctx, radix.Cmd(&str, "ECHO", val)); err != nil {
				t.Errorf("Error al ejecutar el comando, %+v", err)
				return
			}
			if str != val {
				t.Errorf("el valor ejecutado %s es diferente al recibido %s", val, str)
			}
		})
	})

}

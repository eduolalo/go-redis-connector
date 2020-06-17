package redis

import (
	"testing"

	"github.com/mediocregopher/radix/v3"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {

	pool := Connect()
	defer pool.Close()

	var str string
	require.Nil(t, pool.Do(radix.Cmd(&str, "ECHO", "jaló esta merga")))
	require.Equal(t, "jaló esta merga", str)
}

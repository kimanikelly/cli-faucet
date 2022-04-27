package contract

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderConnection(t *testing.T) {

	ethClientType := fmt.Sprintf("%T", ProviderConnection())

	assert.Equal(t, ethClientType, "*ethclient.Client", "The ProviderConnection function should return the type *ethclient.Client")

}

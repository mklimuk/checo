package lichess

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_GetAccount(t *testing.T) {
	c := NewClient("mslG0POzpfuqYIh3", "https://lichess.org")
	acc, err := c.GetAccount()
	require.NoError(t, err)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	require.NoError(t, enc.Encode(acc))
}


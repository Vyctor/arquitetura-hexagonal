package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(test *testing.T) {
	message := "Hello Json"
	result := jsonError(message)
	require.Equal(test, []byte(`{"message":"Hello Json"}`), result)
}

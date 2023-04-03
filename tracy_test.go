package ehtracygo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	correlationID string = "0123456789"
)

func Test_CorrelationID(t *testing.T) {
	ctx := NewContext(context.Background(), correlationID)

	corrID := FromContext(ctx)

	require.Equal(t, correlationID, corrID)
}

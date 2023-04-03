package ehtracygo

import (
	"context"

	eh "github.com/looplab/eventhorizon"
)

// Strings used to marshal context values.
const (
	DefaultCorrelationID string = ""
	correlationIDKeyStr  string = "X-Correlation-ID"
)

func init() {
	eh.RegisterContextMarshaler(func(ctx context.Context, vals map[string]interface{}) {
		if correlationID, ok := ctx.Value(correlationIDKeyStr).(string); ok {
			vals[correlationIDKeyStr] = correlationID
		}
	})

	eh.RegisterContextUnmarshaler(func(ctx context.Context, vals map[string]interface{}) context.Context {
		if correlationID, ok := vals[correlationIDKeyStr].(string); ok {
			ctx = NewContext(ctx, correlationID)
		}

		return ctx
	})
}

// FromContext returns the correlationID from the context, or the an empty string.
func FromContext(ctx context.Context) string {
	if correlationID, ok := ctx.Value(correlationIDKeyStr).(string); ok {
		return correlationID
	}

	return DefaultCorrelationID
}

// NewContext sets the correlationID to use in the context.
func NewContext(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIDKeyStr, correlationID)
}

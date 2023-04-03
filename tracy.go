package ehtracygo

import (
	"context"

	eh "github.com/looplab/eventhorizon"
)

// Strings used to marshal context values.
const (
	correlationIDKeyStr string = "X-Correlation-ID"
)

func init() {
	eh.RegisterContextMarshaler(func(ctx context.Context, vals map[string]interface{}) {
		if ns, ok := ctx.Value(correlationIDKey).(string); ok {
			vals[correlationIDKeyStr] = ns
		}
	})
	eh.RegisterContextUnmarshaler(func(ctx context.Context, vals map[string]interface{}) context.Context {
		if ns, ok := vals[correlationIDKeyStr].(string); ok {
			ctx = NewContext(ctx, ns)
		}

		return ctx
	})
}

type contextKey int

const (
	correlationIDKey contextKey = iota
)

// FromContext returns the correlationID from the context, or the default correlationID.
func FromContext(ctx context.Context) string {
	if ns, ok := ctx.Value(correlationIDKey).(string); ok {
		return ns
	}

	return ""
}

// NewContext sets the correlationID to use in the context. The correlationID is used to
// determine which database to use, among other things.
func NewContext(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIDKey, correlationID)
}

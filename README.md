# Info

This libary can be used to enhance the https://github.com/looplab/eventhorizon library with the functionality to trace the `X-Correlation-ID` through events.

Supported Go Versions:

This library supports two most recent Go, currently 1.16

# Install

```bash
go get github.com/Clarilab/eh-tracygo
```

# Usage

Get Correlation-ID from Context:

```Go
import (
    ehtracygo "github.com/Clarilab/eh-tracygo"
)

func someFunction(ctx context.Context) {
    correlationID := ehtracygo.FromContext(ctx)
}

```

Add Correlation-ID to Context:

```Go
import (
    ehtracygo "github.com/Clarilab/eh-tracygo"
)

func someFunction(ctx context.Context, correlationID string) {
    ctx = ehtracygo.NewContext(ctx, correlationID)
}

```
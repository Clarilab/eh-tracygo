# Info

This libary can be used to enhance the [looplab/eventhorizon](https://github.com/looplab/eventhorizon) library with the functionality to trace the `X-Correlation-ID` through events.

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

If the Correlation-ID is already present in the context, simply initialize this libary to use it in combination with the [looplab/eventhorizon](https://github.com/looplab/eventhorizon) library.

```Go
import (
    _ "github.com/Clarilab/eh-tracygo"
)
```
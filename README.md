# hello otel

Small example using opentelemetry instrumentation and jaeger for visualization. Just 2 identical binaries ping-ponging each other with http requests and a random sleep, nothing fancy.

Just do `docker compose up --build` followed by `go test -bench-.` to get some data.

You can then access `http://localhost:16686` for the jaeger dashboard.
package internal

import (
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func ping(pongerAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pongerAddr, err := url.Parse(pongerAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := otelhttp.Get(r.Context(), fmt.Sprintf("%s/pong", pongerAddr.String()))
		if err != nil {
			w.WriteHeader(res.StatusCode)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		defer res.Body.Close()

		pong, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(pong)
	}
}

func pong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wait := rand.Int31n(50)

		time.Sleep(time.Millisecond * time.Duration(wait))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}
}

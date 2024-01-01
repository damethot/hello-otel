package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func BenchmarkPingPong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err := http.Get("http://localhost:3000/ping")
		if err != nil {
			b.FailNow()
		}
		defer res.Body.Close()

		pong, _ := io.ReadAll(res.Body)

		fmt.Println(string(pong))
	}
}

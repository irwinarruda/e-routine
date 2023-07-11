package main

import (
	"e-routine/shared/http"
	"e-routine/shared/providers/env"

	_ "github.com/lib/pq"
)

func main() {
	env.Load()
	http.Start()
}

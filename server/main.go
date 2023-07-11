package main

import (
	"e-routine/shared/http"
	"e-routine/shared/providers/env"
)

func main() {
	env.Load()
	http.Start()
}

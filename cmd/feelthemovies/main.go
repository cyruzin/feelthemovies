package main

import (
	"context"

	"github.com/cyruzin/feelthemovies/internal/app/server"
	"github.com/cyruzin/feelthemovies/internal/pkg/ratelimit"
)

// Runs a background goroutine to remove
// old entries from the visitors map.
func init() {
	go ratelimit.CleanupVisitors()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.New(ctx)
}

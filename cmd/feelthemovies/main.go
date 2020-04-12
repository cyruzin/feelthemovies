package main

import (
	"context"

	"github.com/cyruzin/feelthemovies/internal/app/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.New(ctx)
}

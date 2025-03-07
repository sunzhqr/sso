package main

import (
	"fmt"
	"github.com/sunzhqr/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println("This is config:", cfg)

	// TODO: initialize logger

	// TODO: initialize app

	// TODO: launch gPRC-server of the app
}

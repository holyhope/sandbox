package run

import (
	"fmt"
	"github.com/holyhope/sandbox/auto-init/v2"
)

// Deprecated: use Run instead.
func OldRun() {
	Run()
}

func Run() {
	if !v2.IsInit() {
		panic("not yet initialized")
	}

	fmt.Println("Running")
}
package main

import (
	v2 "github.com/holyhope/sandbox/auto-init/v2"
	"github.com/holyhope/sandbox/run"
)

func main() {
	v2.Init()
	run.Run()
}

package v2

import "fmt"

var initialized bool

func Init() {
	fmt.Println("Initialized")

	initialized = true
}

func IsInit() bool {
	return initialized
}

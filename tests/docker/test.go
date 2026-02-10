package main

import (
	"fmt"
	"runtime"

	"github.com/Hayao0819/go-distro/linux"
)

func main() {
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Println()

	os := linux.Get()
	fmt.Printf("ID: %s\n", os.ID())
	fmt.Printf("FullName: %s\n", os.FullName())
	fmt.Printf("VerID: %s\n", os.VerID())
	fmt.Printf("VerFullName: %s\n", os.VerFullName())
	fmt.Printf("VerCodeName: %s\n", os.VerCodeName())
}

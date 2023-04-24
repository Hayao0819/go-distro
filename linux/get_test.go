package linux_test

import (
	"fmt"
	"testing"

	"github.com/Hayao0819/go-distro/linux"
)

func TestGet(t *testing.T) {
	l := linux.Get()
	fmt.Println(l.ID())
}

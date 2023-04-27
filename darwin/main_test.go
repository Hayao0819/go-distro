package darwin_test

import (
	"fmt"
	"testing"

	"github.com/Hayao0819/go-distro/darwin"
)

func TestGet(t *testing.T) {
	d := darwin.Get()
	fmt.Println(d.ID())
}

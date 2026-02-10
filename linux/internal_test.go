package linux

import (
	"testing"

	"github.com/Hayao0819/go-distro/base"
	"github.com/stretchr/testify/assert"
)

func TestConstValue(t *testing.T) {
	testcase := []struct {
		c *Linux
		s base.ID
	}{
		{
			c: Arch,
			s: "arch",
		},
		{
			c: CentOS,
			s: "centos",
		},
		{
			c: Other,
			s: "other",
		},
	}

	for _, test := range testcase {
		assert.Equal(t, test.c.ID(), test.s)
	}
}

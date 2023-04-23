package linux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstValue(t *testing.T) {
	testcase := []struct {
		c *Linux
		s string
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
		assert.Equal(t, test.c.String(), test.s)
	}
}



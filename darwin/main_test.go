package darwin

import (
	"testing"

	"github.com/Hayao0819/go-distro/base"
)

func TestGetFromVersion(t *testing.T) {
	testcase := []struct {
		verstr   string
		expectid base.ID
	}{
		{
			verstr:   "9.0",
			expectid: Other.VerID(),
		},
		{
			verstr:   "10.0",
			expectid: Cheetah.VerID(),
		},
		{
			verstr:   "10.1",
			expectid: Puma.VerID(),
		},
		{
			verstr:   "10.2.1",
			expectid: Jaguar.VerID(),
		},
		{
			verstr:   "10.3",
			expectid: Panther.VerID(),
		},
		{
			verstr:   "10.4",
			expectid: Tiger.VerID(),
		}, {
			verstr:   "10.5",
			expectid: Leopard.VerID(),
		},
		{
			verstr:   "10.6",
			expectid: SnowLeopard.VerID(),
		},
		{
			verstr:   "10.7",
			expectid: Lion.VerID(),
		},
		{
			verstr:   "10.8",
			expectid: MountainLion.VerID(),
		},
		{
			verstr:   "10.9",
			expectid: Mavericks.VerID(),
		},
		{
			verstr:   "10.10",
			expectid: Yosemite.VerID(),
		},
		{
			verstr:   "10.11",
			expectid: ElCapitan.VerID(),
		},
		{
			verstr:   "10.12",
			expectid: Sierra.VerID(),
		},
		{
			verstr:   "10.13",
			expectid: HighSierra.VerID(),
		},
		{
			verstr:   "10.14",
			expectid: Mojave.VerID(),
		},
		{
			verstr:   "10.15",
			expectid: Catalina.VerID(),
		},
		{
			verstr:   "10.16",
			expectid: BigSur.VerID(),
		},
		{
			verstr:   "11.0",
			expectid: BigSur.VerID(),
		},
		{
			verstr:   "12.0",
			expectid: Monterey.VerID(),
		},
		{
			verstr:   "13.3",
			expectid: Ventura.VerID(),
		},
		{
			verstr:   "14.0",
			expectid: Sonoma.VerID(),
		},
		{
			verstr:   "14.5",
			expectid: Sonoma.VerID(),
		},
		{
			verstr:   "15.0",
			expectid: Sequoia.VerID(),
		},
		{
			verstr:   "15.2",
			expectid: Sequoia.VerID(),
		},
		{
			verstr:   "26.0",
			expectid: Tahoe.VerID(),
		},
	}

	for _, tt := range testcase {
		if getFromVersion(tt.verstr).VerID() != tt.expectid {
			t.Errorf("got %s, want %s", getFromVersion(tt.verstr).VerID(), tt.expectid)
		}
	}
}

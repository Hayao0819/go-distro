package darwin

import (
	"testing"
)

func TestGetFromVersion(t *testing.T){
	testcase := []struct{
		verstr string
		expectid string
	}{
		{
			verstr: "9.0",
			expectid: Other.Version().ID(),
		},
		{
			verstr: "10.0",
			expectid: Cheetah.Version().ID(),
		},
		{
			verstr: "10.1",
			expectid: Puma.Version().ID(),
		},
		{
			verstr: "10.2.1",
			expectid: Jaguar.Version().ID(),
		},
		{
			verstr: "10.3",
			expectid: Panther.Version().ID(),
		},
		{
			verstr: "10.4",
			expectid: Tiger.Version().ID(),
		},{
			verstr: "10.5",
			expectid: Leopard.Version().ID(),
		},
		{
			verstr: "10.6",
			expectid: ShowLeopard.Version().ID(),
		},
		{
			verstr: "10.7",
			expectid: Lion.Version().ID(),
		},
		{
			verstr: "10.8",
			expectid: MountainLion.Version().ID(),
		},
		{
			verstr: "10.9",
			expectid: Mavericks.Version().ID(),
		},
		{
			verstr: "10.10",
			expectid: Yosemite.Version().ID(),
		},
		{
			verstr: "10.11",
			expectid: ElCapitan.Version().ID(),
		},
		{
			verstr: "10.12",
			expectid: Serra.Version().ID(),
		},
		{
			verstr: "10.13",
			expectid: HighSerra.Version().ID(),
		},
		{
			verstr: "10.14",
			expectid: Mojave.Version().ID(),
		},
		{
			verstr: "10.15",
			expectid: Catalina.Version().ID(),
		},
		{
			verstr: "10.16",
			expectid: BigSur.Version().ID(),
		},
		{
			verstr: "11.0",
			expectid: BigSur.Version().ID(),
		},
		{
			verstr: "12.0",
			expectid: Monterey.Version().ID(),
		},
		{
			verstr: "13.3",
			expectid: Ventura.Version().ID(),
		},
	}

	for _, tt := range testcase{
		if getFromVersion(tt.verstr).Version().ID() != tt.expectid{
			t.Errorf("got %s, want %s", getFromVersion(tt.verstr).Version().ID(), tt.expectid)
		}
	}
}

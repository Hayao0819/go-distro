package lsbrelease_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Hayao0819/go-distro/lsbrelease"
)

func TestRead(t *testing.T) {
	type lsbjson struct {
		ID          string `json:"id"`
		DESCRIPTION string `json:"desc"`
		RELEASE     string `json:"release"`
		CODENAME    string `json:"codename"`
	}

	l, err := lsbrelease.Read()
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(l, "", "  ")
	fmt.Println(string(b))

}

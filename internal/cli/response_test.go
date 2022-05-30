package cli_test

import (
	"os"
	"testing"

	"github.com/pauloalchemist/climabr/internal/cli"
)

// Test ParseResponse with api data in a JSON file
func TestParseResponse(t *testing.T) {
	data, err := os.ReadFile("../../.dados/fsa_advisor_response.json")
	cli.CheckError(err)

	expect := cli.Conditions{
		City: "Feira de Santana",
	}

	got, err := cli.ParseResponse(data)
	cli.CheckError(err)

	if expect.City != got.City {
		t.Errorf("Expect city: %q | But got city: %q.  ", expect, got.City)
	}
}

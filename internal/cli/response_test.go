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
		Temp: 21,
	}

	got, err := cli.ParseResponse(data)
	cli.CheckError(err)

	t.Run("ParseExpectCity", func(t *testing.T) {
		if expect.City != got.City {
			t.Errorf("Expect city: %v | But got city: %v.  ", expect.City, got.City)
		}
	})

	t.Run("ParseExpectTemp", func(t *testing.T) {
		if expect.Temp != got.Temp {
			t.Errorf("Expect temp: %v | But got temp: %v.  ", expect.Temp, got.Temp)
		}
	})
}

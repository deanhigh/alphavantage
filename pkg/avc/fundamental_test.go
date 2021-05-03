package alphavantage

import (
	"encoding/json"
	"testing"
)

// TestCompanyOverview Testing the parsing of company overview returned by the api
func TestCompanyOverview(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	co, err := c.FundamentalService.GetCompanyOverview("TDOC")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Validating company overview: \n%#v", co)
	if co.Symbol != "TDOC" {
		t.Errorf("CompanyOverview = %v; looking for TDOC", co)
	}

	cosJSON, _ := json.Marshal(co)
	t.Logf("%s", cosJSON)
}

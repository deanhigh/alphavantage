package alphavantage

import (
	"testing"
)

// GetRequestTest Testing the parsing of company overview returned by the api
func TestGetRequest(t *testing.T) {
	co := CompanyOverview{}
	c, err := NewClient()

	if err != nil {
		t.FailNow()
	}

	params := map[string]string{"function": "OVERVIEW", "symbol": "IBM"}

	req, err := c.NewQuery(params)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Do(req, &co)
	if err != nil {
		// io.Copy(os.Stdout, resp.Body)
		t.Logf("%3v", resp.Body)
		t.Fatal(err)
	}

	// t.Logf("Response: \n%#v", resp)
	t.Logf("Validating company overview: \n%#v", co)
	if co.Symbol != "IBM" {
		t.Errorf("CompanyOverview = %v; looking for APPL", co)
	}

}

package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type Sample struct {
	SampleString        string           `json:"SampleString"`
	SampleInt           int              `json:"SampleInt"`
	SampleOptionalDate  *OptionalDate    `json:"SampleOptionalDate,string"`
	SampleOptionalFloat *OptionalFloat64 `json:"SampleOptionalFloat,string"`
}

type TransformSample struct {
	SampleString        string     `json:"SampleString"`
	SampleInt           int        `json:"SampleInt"`
	SampleOptionalDate  *time.Time `json:"SampleOptionalDate"`
	SampleOptionalFloat float64    `json:"SampleOptionalFloat,float64"`
}

// GetRequestTest Testing the parsing of company overview returned by the api
func TestDumpJSON(t *testing.T) {
	testTime := time.Now()
	sod := OptionalDate(testTime)
	f := float64(16.2)
	sof := OptionalFloat64{value: &f}
	s := Sample{SampleString: "Test", SampleInt: 5, SampleOptionalDate: &sod, SampleOptionalFloat: &sof}
	s2 := TransformSample{SampleString: "Test", SampleInt: 5, SampleOptionalDate: &testTime, SampleOptionalFloat: 16.2}

	fo, err := ioutil.TempFile("", "dumptest.*.json")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(fo.Name())

	DumpJSON(fo, s)

	fo.Close()

	data, err := ioutil.ReadFile(fo.Name())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", string(data))

	v := TransformSample{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(v, s2); diff != "" {
		t.Errorf("TransformCompanyOverview hasn't saved properly to disk, \n%s", cmp.Diff(v, s2))
	}

}

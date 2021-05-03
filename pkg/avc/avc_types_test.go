package alphavantage

import (
	"encoding/json"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%#v != %#v", a, b)
	}
}

// TestOptionalFloat AlphaVantage presents floats as string in Json so that it can use None if they explicitly don't want to set a value
// these tests test the OptionalFloat construct which loads those values and ignores the None values.
func TestOptionalFloat(t *testing.T) {
	type FakeWrapper struct {
		FakeFloat *OptionalFloat64 `json:"fakeFloat,omitempty"`
	}
	t.Run("TestOptionalFloat unmarshal basic float value", func(t *testing.T) {
		sampleJSON := `{"fakeFloat":"0.1234"}`
		fakeWrapper := FakeWrapper{}
		json.Unmarshal([]byte(sampleJSON), &fakeWrapper)
		assertEqual(t, float64(*fakeWrapper.FakeFloat.value), 0.1234)
	})
	t.Run("TestOptionalFloat unmarshal None value", func(t *testing.T) {
		sampleJSON := `{"fakeFloat":"None"}`
		fakeWrapper := FakeWrapper{}
		json.Unmarshal([]byte(sampleJSON), &fakeWrapper)
		if fakeWrapper.FakeFloat.value != nil {
			t.Fatalf("Float value expected to be nil, found %#v", fakeWrapper.FakeFloat.value)
		}
	})
}

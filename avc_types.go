package alphavantage

import (
	"strconv"
	"strings"
	"time"
)

// OptionalDate is the wrapper date for alphavantage dates in the api
type OptionalDate time.Time

// OptionalFloat64 is used to handle the fact that AV API can return None or a float64.
type OptionalFloat64 struct {
	value *float64
}

// UnmarshalJSON Unmarshal optional types from the AV api which can include None or a float64 number.
func (j *OptionalFloat64) UnmarshalJSON(b []byte) error {
	// If None, return nil because AV api allows strings of None to be returned
	s := strings.Trim(string(b), "\"")
	if s == "None" {
		// Don't do anything and return
		*j = OptionalFloat64{value: nil}
		return nil
	}
	// If not None, then try parse it.
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*j = OptionalFloat64{value: &val}

	return nil
}

// MarshalJSON marshals the OptionalFloat64 type into float64 output fot json to avoid wrapping in strings
func (j *OptionalFloat64) MarshalJSON() ([]byte, error) {
	f := j.value
	if f != nil {
		vs := strconv.FormatFloat(float64(*f), 'f', 2, 64)
		return []byte(vs), nil
	}
	return []byte("null"), nil
}

// UnmarshalJSON Unmarsal optional date from AV
func (j *OptionalDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "None" {
		// Don't do anything and return
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = OptionalDate(t)
	return nil
}

// MarshalJSON marshal a jsondate
func (j OptionalDate) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	return t.MarshalJSON()
}

// GoString for printing your date to %v
func (j OptionalDate) GoString() string {
	t := time.Time(j)
	return t.Format("2006-01-02")
}

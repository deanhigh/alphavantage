package alphavantage

import (
	"encoding/json"
	"os"
)

// DumpJSON dumps the alphavantage objects as is.
func DumpJSON(fo *os.File, item interface{}) error {
	data, err := json.MarshalIndent(item, "", " ")
	if err != nil {
		return err
	}

	_, err = fo.Write(data)
	return err
}

package valr

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

func PrettyPrint(jString string) {
	PrettyPrintBytes([]byte(jString))
}

func PrettyPrintBytes(jBytes []byte) (string, error) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, jBytes, "", "\t")
	if error != nil {
		return "", error
	}

	return string(prettyJSON.Bytes()), nil
}

type quotedInt int32

func (d *quotedInt) UnmarshalJSON(data []byte) error {
	res, err := strconv.Atoi(strings.Trim(string(data), `"'`))
	if err != nil {
		return err
	}

	*d = quotedInt(res)

	return nil
}

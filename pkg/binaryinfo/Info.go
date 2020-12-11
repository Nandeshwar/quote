package binaryinfo

import (
	"bytes"
	"encoding/json"
)

type InfoMohan struct {
	Id        uint8
	Name      [10]byte // max 10 bytes
	Fee       float64
	CreatedAt [25]byte // 2012-11-01T22:08:41+00:00 (25 char length) - RFC3339
}

const INFO_TYPE = 1

const INFO_LEN = 44

//const INFO_LEN = 11

func (i *InfoMohan) ToJSON() (string, error) {

	jsonBytes, err := i.MarshalJSON()
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// Custom JSON Marshalling to output time field in required format
// http://choly.ca/post/go-json-marshalling/
func (d *InfoMohan) MarshalJSON() ([]byte, error) {
	trimmedBytesName := trimNullInBytesArray(d.Name[:])
	trimmedBytesCreatedAt := trimNullInBytesArray(d.CreatedAt[:])
	type Alias InfoMohan
	return json.Marshal(&struct {
		*Alias
		Name      string `json:"Name"`
		CreatedAt string `json:"CreatedAt"`
	}{
		Alias: (*Alias)(d),
		//Name:  string(d.Name[:]),
		Name:      string(trimmedBytesName[:]),
		CreatedAt: string(trimmedBytesCreatedAt[:]),
	})
}

func trimNullInBytesArray(bArr []byte) []byte {
	return bytes.Trim(bArr, string([]byte{0x00}))
}

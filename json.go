package main

import (
	"bytes"
	"encoding/json"
)

func (n Name) ComposeNameJSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(name)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ParseNameJSON(b []byte) (Name, error) {
	n := Name{}
	r := bytes.NewReader(b)
	err := json.NewDecoder(r).Decode(&n)
	return n, err
}

package httpmongo

import (
	"encoding/json"
)

func insert_jsonstring(DB string, C string, jsonstr string) error {
	return insert_jsonbytes(DB, C, []byte(jsonstr))
}

func insert_jsonbytes(DB string, C string, jsonbytes []byte) error {
	c := MgoDataCollect(DB, C)

	var document interface{}
	err := json.Unmarshal(jsonbytes, &document)
	if err != nil {
		return err
	}
	err = c.Insert(document)
	if err != nil {
		return err
	}
	return nil
}


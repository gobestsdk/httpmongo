package httpmongo

import (
	"encoding/json"
	"log"
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

func find_filterstring(DB string, C string, filterstr string) []JsonDocument {
	c := MgoDataCollect(DB, C)
	var filter interface{}
	err := json.Unmarshal([]byte(filterstr), &filter)
	if err != nil {
		panic(err)
	}
	log.Print(filter)
	query := c.Find(nil)

	result := []JsonDocument{}

	query.All(&result)
	return result
}

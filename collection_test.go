package httpmongo

import (
	"log"
	"encoding/json"
	"testing"
)

func Test_One(t *testing.T){
	var str ="[{\"name\":\"l1\"},{\"name\":\"l2\"}]"
	var wo []map[string]string
	json.Unmarshal([]byte(str),&wo)
	log.Print(str)
	log.Print(wo)
}
package httpmongo

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/golangframework/File"
)

func Test_One(t *testing.T) {
	var str = "[{\"name\":\"l1\"},{\"name\":\"l2\"}]"
	var wo []map[string]string
	json.Unmarshal([]byte(str), &wo)
	log.Print(str)
	log.Print(wo)
}

func Test_update(t *testing.T) {
	var updatestr = File.ReadAllText("./update.json")
	update("DB")
}

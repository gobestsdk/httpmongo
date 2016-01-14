// urlparser_tesg.go
package httpmongo

import (
	"fmt"
	"testing"
)

func Test_Mongo_parse(t *testing.T) {

	cmd, _ := Mongo_parse("/mongo.show dbs")
	fmt.Printf(cmd)
}
func Test_Mongo_DB_parse(t *testing.T) {

	cmd, DB, _ := Mongo_DB_parse("/mongo/Insurance.show collections")
	fmt.Printf(cmd + "\t" + DB)
}
func Test_Mongo_DB_C_parse(t *testing.T) {

	cmd, DB, C, _ := Mongo_DB_C_parse("/mongo/Insurance/Vehicle.insert(" + "{}" + ")")
	fmt.Printf(cmd + "\t" + DB + "\t" + C)
}

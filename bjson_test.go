package bjson

import (
	"testing"
	"log"
)

func TestBjson_MapString(t *testing.T) {
	bs := []byte(`{"name":"bysir","sex":1,"age":21,"data":{"hab":"code"}}`)
	bj, _ := New(bs)

	ms := bj.MapString()
	mi := bj.MapInterface()

	log.Print("ms: ", ms)
	log.Print("mi: ", mi)
	log.Printf("name: %s", bj.Pos("name").String()) //bysir
	log.Printf("age: %d,%s", bj.Pos("age").Int(), bj.Pos("age").String()) // 21,21
	log.Printf("sex: %t,%d", bj.Pos("sex").Bool(), bj.Pos("sex").Int()) // true,1
	log.Printf("hab: %s", bj.Pos("data").Pos("hab").String()) // code

	log.Printf("E name: %d", bj.Pos("name").Int()) // 0

}

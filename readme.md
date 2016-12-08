# bjson
golang json util 

## show me the code 
bjson_test.go
```go
bs := []byte(`{"name":"bysir","sex":1,"age":21,"data":{"hab":"code"},"extra":["a",1]}`)
bj, _ := New(bs)

log.Print("ms: ", bj.MapString()) // map[extra: name:bysir sex:1 age:21 data:]
log.Print("mi: ",  bj.MapInterface()) //  map[name:bysir sex:1 age:21 data:map[hab:code] extra:[a 1]]
log.Printf("name: %s", bj.Pos("name").String()) // bysir
log.Printf("age: %d,%s", bj.Pos("age").Int(), bj.Pos("age").String()) // 21,21
log.Printf("sex: %t,%d", bj.Pos("sex").Bool(), bj.Pos("sex").Int()) // true,1
log.Printf("hab: %s", bj.Pos("data").Pos("hab").String()) // code
log.Printf("ext: len:%d data[0]:%s",bj.Pos("extra").Len(), bj.Pos("extra").Index(0).String()) // 2,a

log.Printf("E name: %d", bj.Pos("name").Int()) // 0
```

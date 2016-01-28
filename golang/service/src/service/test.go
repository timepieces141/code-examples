package main

import (
	"encoding/json"
	"fmt"
)

var data string = `
{
	"foo":"bar"
}
`
type Foo struct {
	Foo string
}


func main() {
	var foo Foo
	err := json.Unmarshal([]byte(data), &foo)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", foo)
}
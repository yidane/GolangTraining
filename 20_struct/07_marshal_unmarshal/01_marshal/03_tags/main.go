package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
	name  string `json:"Name"`
}

func main() {
	p1 := person{"James", "Bond", 20, "yidane"}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

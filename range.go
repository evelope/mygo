package main

import (
	"fmt"
)

func main() {
	nums:=[]int{1,3,5}

	for key,num:=range nums {
		fmt.Println(key,num)
	}

	kvs:=map[string]string{"a": "aa","b": "bb"}
	for key,kv:=range kvs {
		fmt.Println(key,kv)
	}
}
package main

import (
	"fmt"
	)

type tst struct {
	list string
	tstmap map[string][]int
}


func main() {

	x := new(tst)
	intslice := make([]int,5)
	x.tstmap = make(map[string][]int)

	x.tstmap["first"] = intslice
	intslice[0] = 10
	fmt.Println("map: ", x.tstmap["first"][0])
}

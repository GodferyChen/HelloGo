package main

import "fmt"

type Ver struct {
	Lat,Long float64
}

var m map[string]Ver

func main() {
	m = make(map[string]Ver)
	m["Bell Labs"] = Ver{
		40.68433,-74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

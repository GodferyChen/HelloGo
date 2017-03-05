package main

import "fmt"

type Vert struct {
	Lat,Long float64
}

var v = map[string]Vert{
	"Bell Labs":{40.68433,-74.39967},
	"Google":{37.4222,23.22122},
}

func main() {
	fmt.Println(v)

}

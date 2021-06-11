package main

import (
	"encoding/json"
	"fmt"
)

// {"name":"martin","loc":{"lat":51.33962,"long":12.37129}}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Peer struct {
	Name     string   `json:"name"`
	Location Location `json:"loc"`
}

func main() {
	peer := Peer{
		Name: "martin",
		Location: Location{
			Lat:  51.33962,
			Long: 12.37129,
		}}
	fmt.Println(peer)
	// {martin {51.33962 12.37129}}
	fmt.Printf("%v\n", peer)
	fmt.Printf("%+v\n", peer)
	fmt.Printf("%#v\n", peer)

	// int64(1.2) // constant 1.2 truncated to integer

	b, _ := json.Marshal(peer)
	fmt.Println(string(b))
}

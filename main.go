package main

import (
	"log"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"gray":  "#898989",
	// }
	// log.Println(colors)

	// person := make(map[string]string)
	// person["name"] = "Mohammed"
	// person["age"] = "40"
	// log.Println(person)

	status := make(map[int]string)
	status[1] = "good"
	status[0] = "bad"
	printMap(status)

}

func printMap(c map[int]string) {
	for number, desc := range c {
		log.Println(number, desc)
	}
}

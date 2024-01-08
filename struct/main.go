package main

import "fmt"

func main() {
	colors := map[string]string{
		"name":   "white",
		"colors": "black",
		"pack":   "red",
	}
	for i, c := range colors {
		fmt.Println(i, c)
	}
}

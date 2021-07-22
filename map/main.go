package main

import "fmt"

func main() {
	colors := get_map_a()
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex code of", key, "is", val)
	}
}

func get_map_a() map[string]string {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}
	return colors
}

func get_map_b() map[string]string {
	var colors map[string]string = map[string]string{}
	colors["red"] = "#ff000"
	colors["green"] = "#4bf745"
	return colors
}

func get_map_c() map[string]string {
	colors := make(map[string]string)
	colors["red"] = "#ff000"
	colors["green"] = "#4bf745"
	delete(colors, "red")
	return colors
}

package main

import "fmt"

func main(){
	/**

		1. Maps are similar to structs in the way that, they too hold collection
			of values as key-value pairs.

		2. The major differnce is that we can iterate over maps and we cannot iterate ove rstructs.

		3. Maps are pointer types whereas structs are value types.
	 */

		colors := map[string]string {
			"red":"#ff0000",
			"green":"#00ff00",
		}

		colors["white"] = "#FFFFFF"
		printMap(colors)
}

func printMap(colors map[string]string) {
	for i,j := range colors {
		fmt.Println(i,j)
	}
}
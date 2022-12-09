package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
	}

	changeMap(m)

	fmt.Println(m)
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
	m["goat"] = "meel"
}

//func changeMaps(m map[string]string) {
//	m["goat"] = "meel"
//}

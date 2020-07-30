package main

import "fmt"

func printMap(m map[string]string) {
	fmt.Printf("m = %v, len(m) = %d\n", m, len(m))
}

func main() {
	m := map[string]string{
		"name":    "xiangnan",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 = map[string]int{}  // m3 == nil
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Printf("k=%s, v=%s\n", k, v)
	}

	fmt.Println("Getting values...")
	courseName, ok := m["course"]
	fmt.Println("courseName:", courseName, ", ok:", ok)
	if courName, ok := m["cour"]; ok {
		fmt.Println("courName:", courName, ", ok:", ok)
	} else {
		fmt.Println("m does not have the key: cour")
	}

	fmt.Println("Deleting elements...")
	printMap(m)
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	printMap(m)

	m["gender"] = "female"
	m["quality"] = "good"
	printMap(m)
}

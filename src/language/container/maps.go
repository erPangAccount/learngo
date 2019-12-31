package main

import "fmt"

func main() {
	fmt.Println("Creating map")
	m := map[string]string{
		"name":    "name",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3) //map[course:golang name:name quality:notbad site:imooc] map[] map[]

	for k, v := range m {
		fmt.Println(k, v)
		//name name
		//course golang
		//site imooc
		//quality notbad
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok) //golang true

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exists") //key does not exists
	}

	fmt.Println("Deleting values")
	fmt.Println(m) //map[course:golang name:name quality:notbad site:imooc]
	delete(m, "name")
	fmt.Println(m) //map[course:golang quality:notbad site:imooc]

}

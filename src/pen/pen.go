package main

import (
	"fmt"
)

func main() {
	for {
		var ans int

		fmt.Println("I have a ?")
		fmt.Println("1: pen\n2: eraser\n3: dream")
		fmt.Scan(&ans)
		if 1 != ans {
			continue
		}

		fmt.Println("I have a ?")
		fmt.Println("1: banana\n2: apple\n3: issue")
		fmt.Scan(&ans)
		if 2 != ans {
			continue
		}
		fmt.Println("apple pen!")

		fmt.Println("I have a ?")
		fmt.Println("1: pen\n2: eraser\n3: dream")
		fmt.Scan(&ans)
		if 1 != ans {
			continue
		}

		fmt.Println("I have a ?")
		fmt.Println("1: pineapple\n2: cherry\n3: plum")
		fmt.Scan(&ans)
		if 1 != ans {
			continue
		}
		fmt.Println("pineapple pen!")

		fmt.Println("apple pen?")
		fmt.Println("1: yes\n2: no\n")
		fmt.Scan(&ans)
		if 1 != ans {
			continue
		}

		fmt.Println("pineapple pen?")
		fmt.Println("1: yes\n2: no\n")
		fmt.Scan(&ans)
		if 1 != ans {
			continue
		}

		fmt.Println("penpineappleapplepen!")
	}
}

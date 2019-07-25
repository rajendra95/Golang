// Fibonaci series *****************************

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	f := 0
	s := 1
	//var count int
	for i := 0; i < 5; i++ {	

		temp := f
		f = s
		s = temp + s
	}
	fmt.Println(f)
}

// reverse the string ****************************
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	s := "meGhaNa"

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	fmt.Println(s)
	fmt.Println(string(runes))
}


//find out duplicates in the string ***************************
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	s := "12345612"
	runes := []rune(s)
	//res := make([]string, 1)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				fmt.Println(string(runes[i]))
			}
		}
		//fmt.Println(string(runes[i]))
	}
}





//how to duplicates replace them in the string ***********************
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	a := []int{1, 2, 3, 4, 5, 6, 9, 2, 3}
	// we will use a map for checking the value is not repeated
	keys := make(map[int]bool)
	// empty slice to append the elements
	res := make([]int,0)
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			res = append(res, entry)
		}
	}
	fmt.Println("resultant array ", res)
}



//sort the elemets of array using bubble sort ***********************

func main() {
	fmt.Println("Hello, playground")
	a := []int{1, 32, 18, 45, 67, 13, 2, 56, 33, 104, 4}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp

			}
		}
	}
	fmt.Println("New array is ", a)

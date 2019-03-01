package main

import "fmt"

func main(){
 	var a = []int{1,2,3,4,5}
	sum:=sumOfnums(a)
	fmt.Println("sum of numbers is %v",sum)
}

func sumOfnums(arr []int)int{
	sum := 0
	fmt.Println(arr)

	for i:=0;i<len(arr);i++{
		sum+=arr[i]
	}
	return sum
}

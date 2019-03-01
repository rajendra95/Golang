package main

import "fmt"


func main(){
	squareChan := make (chan int)
	cubeChan := make(chan int)
	go squarenum(squareChan)
	go cubenum(cubeChan)
	// its time to send an integer on the channel
	num:= 4
	squareChan <- num
	fmt.Println("squarechan go routine activated")

	cubeChan <-num
	fmt.Println("cubechan go routine activated")

	//its time to read the channels back for the output
	square:= <- squareChan
	cube := <- cubeChan
	totalsum := square+cube
	fmt.Println("the sum is ",totalsum)
}

func squarenum( c chan int){
	fmt.Println("square reading ...")
	num:= <-c    // channel has been read ..
	c<-num*num // writing the square of num into the channel
}

func cubenum(c chan int){
	fmt.Println("cube is reading....")
	num:= <-c  // channel has been read ...
	c <- num*num*num // writing the cube of num into the channel
}

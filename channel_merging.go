package main

import (
	"fmt"
	"time"
	"log"
)

func main(){

	a := channel1(1,2,3,4,5)
	b:= channel1(6,7,8,9,0)
	c:= merge(a,b) // this will merge two channels a & b
	for v:= range c{
		fmt.Println(v)
	}
}

func merge(a,b <- chan int) <- chan int{
	c := make (chan int)
	go func(){
		defer close(c)
		for a!=nil||b!=nil{
			select {
			case v,ok:= <-a:
				if !ok{
					log.Printf("a is done")
					a=nil
					continue
				}
				c<-v
			case v,ok:= <-b:
				if !ok{
					log.Printf("b is done")
					b=nil
					continue
				}
				c<-v
			}
		}
		//close(c)
	}()
return c
}

func channel1(vs...int)<- chan int{
	c:=make(chan int)
	go func(){
		for  _,v:= range vs{
			time.Sleep(time.Second*1)
			c <-v
		}
		close(c)  // after the channl operation is done we will close the channel
	}()
	return c
}

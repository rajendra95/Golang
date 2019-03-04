package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	a:=channelinput(0,1,2,3,4,5,6,7,8,9)
	b:=channelinput(10,11,12,1314,15,16,17,18,19)
	c:=channelinput(20,21,22,23,24,25,26,27,28,29)

	for v:= range mergechannel(a,b,c){
		fmt.Println(v)
	}
}

// one channel for output
//one more channel for input which handles the input from all channels
//once the inputs channels is close the out channel should stop and go routine as well

func mergechannel (chans...<-chan int)<-chan int{
	out:= make (chan int)
	go func(){
		var wg sync.WaitGroup
		wg.Add(len(chans))
		defer close(out)
		for _,c :=range chans{
			go func(c <-chan int){
				for v:=range c{
					out<-v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
func channelinput(vs...int)<- chan int{
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

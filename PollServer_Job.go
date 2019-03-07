package main

import (
	"log"
	"net/http"
	"time"
)

// We have a JO interface which exposes Log method. an interface can pass to this method with a variable number of arguments.
//it also has Suspend and Resume methods.
type Job interface {
	Log(...interface{})
	Suspend() error  // methods implemented by the interface Job
	Resume() error
	Run() error
}

type PollerJob struct {
	suspend chan bool
	resume chan bool
	resourceUrl string
	inMemLog string
} // poller Job struct.

// we will create a newpoller job - it will create a PollerJob with resourceUrl,Suspend,Resume channel for every polling request
func NewPollerJob(resourceUrl string) PollerJob{
	return PollerJob{
		resourceUrl:resourceUrl,
		suspend: make(chan bool),
		resume: make(chan bool),
	}
}
// now we will implement the methods mentioned inside the interface - Job

func (p PollerJob)Log(args...interface{}){
	log.Println(args...)
}

func (p PollerJob)Suspend()error{
	p.suspend <- true
	return nil
}// whenever Suspend method will be called it will write True on the suspend channel and return

func ( p PollerJob) Resume()error{
	p.resume<-true
	return nil
} // whenever resume method will be called it will write true on the resume channel and return

func (p PollerJob)Pollserver()error{
	resp,err:=http.Get(p.resourceUrl)
	if err!=nil{
		return  err
	}
	p.Log(p.resourceUrl,"---",resp.Status)
	return nil
}// whenerver Pollserver Method will be called , it will poll the URL by http.Get method and if there is no error it will print the status

func (p PollerJob)Run()error{
	for{
		select {
		case <-p.suspend:
			<-p.resume
		default:
			if err:=p.Pollserver();err!=nil{
				p.Log("Error trying to fetch the URL",err)
			}
			time.Sleep(1*time.Second)

		}
	}
}

func main(){
	p:=NewPollerJob("https://www.google.com/")
	go p.Run()
	time.Sleep(5*time.Second)

	p.Log("suspending the server for 5 seconds ")
	p.Suspend()
	time.Sleep(5*time.Second)

	p.Log("Resuming Job")
	p.Resume()

	//wait for go routine to complete
	time.Sleep(5* time.Second)

}

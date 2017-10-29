package main

import (
	"fmt"
	"strings"
)


type Handler interface {
	handleRequest(request string)
	setSuccessor(next Handler)
}

type ConcreteHandler1 struct {
	Successor1 Handler
}

type ConcreteHandler2 struct {
	Successor2 Handler
}

type ConcreteHandler3 struct {
	Successor3 Handler
}

func ( *ConcreteHandler1) HandleRequest(request string){
	fmt.Println("R1 got the request...")

	if strings.Compare(request,"R1")==0  {
		fmt.Println( "ConcreteHandler1 => This one is mine!")
	} else {
		var c = new (ConcreteHandler2)
		c.HandleRequest(request)

	}
}

func ( *ConcreteHandler2) HandleRequest(request string){
	fmt.Println("R2 got the request...")
	if strings.Compare(request,"R2")==0  {
		fmt.Println( "ConcreteHandler2 => This one is mine!")
	} else {
		var c = new (ConcreteHandler3)
		c.HandleRequest(request)

	}
}

func ( *ConcreteHandler3) HandleRequest(request string){
	fmt.Println("R3 got the request...")
	if strings.Compare(request,"R3")==0 {
		fmt.Println( "ConcreteHandler3 => This one is mine!")
	}
}


func main(){
	s1 := new (ConcreteHandler1)

	fmt.Println("Sending R2...");
	s1.HandleRequest("R2")
	fmt.Println("Sending R3...");
	s1.HandleRequest("R3")
	fmt.Println("Sending R1...");
	s1.HandleRequest("R1")
	fmt.Println("Sending RX...");
	s1.HandleRequest("RX")


}

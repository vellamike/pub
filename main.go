package main

import "fmt"
import "math/rand"


type pubgoer  struct{
	drinklevel float32
	drinkspeed float32
}

type group  struct{
	members []pubgoer
	turn_to_buy int //index of the member whose turn it is to buy a round
}

type bartender  struct {
	serving *pubgoer
}

type pub  struct {
	groups []*group
}

func main (){	
	fmt.Println("Let's simulate a bar")


	// let's instantiate a group, one person will buy a round, at which
	// point everyone's drink level starts dropping,
	// When the person whose turn it is to buy a round has his drink level drop to zero
	// they offer to buy a round for everyone else,
	// At this point everyone whose glass isn't too full (0.2?) accepts, an order is generated and he goes
	// on the bar queue

	//make a group:#
	pubgoers := make([]pubgoer, 10)
	for i:= 0; i < 10; i++ {
		pubgoers[i] = pubgoer{1.0, rand.Float32()}
	}
}

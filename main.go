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
	serving pubgoer
}

type pub  struct {
	groups []group
	bartenders []bartender
}

func (this *pubgoer) sip() {
	this.drinklevel = this.drinklevel - this.drinkspeed
}

func main (){
	fmt.Println("Let's simulate a bar")


	// let's instantiate a group, one person will buy a round, at which
	// point everyone's drink level starts dropping,
	// When the person whose turn it is to buy a round has his drink level drop to zero
	// they offer to buy a round for everyone else,
	// At this point everyone whose glass isn't too full (0.2?) accepts,
	// an order is generated and goes on the bar queue for a bartender to process when they're
	// available

	//make a group:
	pubgoers := make([]pubgoer, 10)
	for i:= 0; i < 10; i++ {
		pubgoers[i] = pubgoer{1.0, rand.Float32()}
	}

	new_group := group{pubgoers,0}
	fmt.Println(new_group)


	// make some bartenders
	bartenders := make([]bartender, 3, 10)
	for i := 0; i < 3; i++ {
		bartenders[i] = bartender{}
	}
	// now make a pub
	new_pub := pub{[]group{new_group}, bartenders}
	fmt.Println(new_pub)

	drinker := pubgoer{0.9, 0.2}
	fmt.Println(drinker.drinklevel)
	drinker.sip()
	fmt.Println(drinker.drinklevel)
}

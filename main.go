package main

import "fmt"

type Driver struct {
	name string
	//the cost represents how far the driver is from the rider
	cost float32
	occupied bool
	//driver vehicle desc in the format: model/brand/year
	vehicle string
	//if occupied, how much farther is the dropoff cost for the current rider
	costBeforeUnoccupied float32
	//if occupied, the cost + costBeforeUnoccupied will give the 'true cost'
}

func main(){
	fmt.Println("Hello from REBU...")
}
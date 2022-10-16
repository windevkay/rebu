package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/windevkay/rebu/services"
	"github.com/windevkay/rebu/structs"
)

func main() {
	var wg sync.WaitGroup
	driversResultChannel := make(chan structs.Driver)
	control := make(chan structs.Control)

	wg.Add(1)
	go services.GetDriversInRange(2.0, &wg, driversResultChannel, control)

	for {
		select {
		case result := <-driversResultChannel:
			fmt.Println(result)
		case done := <-control:
			if done == structs.Control(0) {
				wg.Wait()
				return
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

package main

import (
	"fmt"
	"github.com/jarvisfriends/go-gpsd"
	"github.com/jarvisfriends/go-gpsd/examples/Console"
)

func main() {
	var gps *gpsd.Session
	var err error

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}
	gps.AddFilter("TPV", Console.OutTPV)
	gps.AddFilter("SKY", Console.OutSKY)
	gps.AddFilter("GST", Console.OutGST)
	gps.AddFilter("ATT", Console.OutATT)
	gps.AddFilter("VERSION", Console.OutVERSION)
	gps.AddFilter("DEVICES", Console.OutDEVICES)
	gps.AddFilter("PPS", Console.OutPPS)
	gps.AddFilter("ERROR", Console.OutERROR)

	gps.AddFilter("TPV", func(r interface{}) {
		tpv := r.(*gpsd.TPVReport)
		fmt.Println("TPV", tpv.Mode, tpv.Time)
	})

	skyFilter := func(r interface{}) {
		sky := r.(*gpsd.SKYReport)

		fmt.Println("SKY", len(sky.Satellites), "satellites")
	}

	gps.AddFilter("SKY", skyFilter)

	gps.AddFilter("TPV", func(r interface{}) {
		t := r.(*gpsd.TPVReport)
		fmt.Println("Received TPV: ", t)

	})
	//Handle errors
	errChan := gps.Watch()
	for {
		if err := <-errChan; err.Error != nil {
			fmt.Println(err.Message)
		} else {
			fmt.Println("OK")
		}
	}
}

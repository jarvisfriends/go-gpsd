package Console

import (
	"fmt"
	"github.com/jarvisfriends/go-gpsd"
)

func OutTPV(r interface{}) {
	t := r.(*gpsd.TPVReport)
	fmt.Println("Received TPV: ", t)
}

func OutSKY(r interface{}) {
	t := r.(*gpsd.SKYReport)
	fmt.Println("Received SKY: ", t)
}

func OutGST(r interface{}) {
	t := r.(*gpsd.GSTReport)
	fmt.Println("Received GST: ", t)
}

func OutATT(r interface{}) {
	t := r.(*gpsd.ATTReport)
	fmt.Println("Received ATT: ", t)
}

func OutVERSION(r interface{}) {
	t := r.(*gpsd.VERSIONReport)
	fmt.Println("Received VERSION: ", t)
}

func OutDEVICES(r interface{}) {
	t := r.(*gpsd.DEVICESReport)
	fmt.Println("Received DEVICES: ", t)
}

func OutPPS(r interface{}) {
	t := r.(*gpsd.PPSReport)
	fmt.Println("Received PPS: ", t)
}

func OutERROR(r interface{}) {
	t := r.(*gpsd.ERRORReport)
	fmt.Println("Received ERROR: ", t)
}

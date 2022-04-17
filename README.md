[![GoDoc](https://godoc.org/github.com/jarvisfriends/go-gpsd?status.svg)](https://godoc.org/github.com/jarvisfriends/go-gpsd)

# go-gpsd

*GPSD client for Go.*

## Installation

<pre><code># go get github.com/jarvisfriends/go-gpsd</code></pre>

go-gpsd has no external dependencies.

## Usage

go-gpsd is a streaming client for GPSD's JSON service and as such can be used only in async manner unlike clients for other languages which support both async and sync modes.
```golang
import ("github.com/jarvisfriends/go-gpsd")

func main() {
	gps := gpsd.Dial("localhost:2947")
}
```

After `Dial`ing the server, you should install stream filters. Stream filters allow you to capture only certain types of GPSD reports.
```golang
gps.AddFilter("TPV", tpvFilter)
```

Filter functions have a type of `gps.Filter` and should receive one argument of type `interface{}`.
```golang
tpvFilter := func(r interface{}) {
	report := r.(*gpsd.TPVReport)
	fmt.Println("Location updated", report.Lat, report.Lon)
}
```

Due to the nature of GPSD reports your filter will manually have to cast the type of the argument it received to a proper `*gpsd.Report` struct pointer.

After installing all needed filters, call the `Watch` method to start observing reports. Please note that at this time installed filters can't be removed.
```golang
errChan := gps.Watch()
```

`Watch()` will span a new goroutine in which all data processing will happen, `errChan` channel will send errors or `nil`. Error handling:

```golang
for {
	if err := <-errChan; err.Error != nil {
		fmt.Println(err.Message)
	} else {
		fmt.Println("OK")
	}
}
```

### Supported GPSD report types

* [`VERSION`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_version) (`gpsd.VERSIONReport`)
* [`TPV`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_tpv) (`gpsd.TPVReport`)
* [`SKY`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_sky) (`gpsd.SKYReport`)
* [`ATT`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_att) (`gpsd.ATTReport`)
* [`GST`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_gst) (`gpsd.GSTReport`)
* [`PPS`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_pps) (`gpsd.PPSReport`)
* [`Devices`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_devices) (`gpsd.DEVICESReport`)
* [`DEVICE`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_device_device) (`gpsd.DEVICEReport`)
* [`ERROR`](https://gpsd.gitlab.io/gpsd/gpsd_json.html#_error) (`gpsd.ERRORReport`)

## Documentation

For complete library docs, visit [GoDoc.org](https://pkg.go.dev/github.com/jarvisfriends/go-gpsd) or take a look at the `gpsd.go` file in this repository.

GPSD's documentation on their JSON protocol can be found at [https://gpsd.gitlab.io/gpsd/gpsd_json.html](https://gpsd.gitlab.io/gpsd/gpsd_json.html)

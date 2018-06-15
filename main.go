package main

import (
	"bytes"
	"encoding/json"
	"flag"
	geo "github.com/rotblauer/google-geolocate"
	"github.com/rotblauer/trackpoints/trackPoint"
	"io"
	"log"
	"net/http"
	"os"

	"fmt"
	"net"
	"time"
)

//https://github.com/skycoin/skycoin/blob/master/src/aether/wifi/wifi.go

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			log.Printf("%v %v\n", i.Name, a)
		}
	}
}

func main() {

	var postUrl string = "DF"
	flag.StringVar(&postUrl, "postUrl", "http://localhost:8080/populate/", "specify where to post")

	flag.Parse()
	client := geo.NewGoogleGeo(os.Getenv("GAPI"))
	point, err := client.Geolocate()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(point)

	fmt.Printf("posting to   %s\n", postUrl)
	t := trackPoint.TrackPoint{Lat: 52.472254, Lng: 13.398756, Time: time.Now()}
	fmt.Print(t)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(t)
	fmt.Print(b)
	result, err := http.Post(postUrl, "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Print(err)
	}
	io.Copy(os.Stdout, result.Body)

}

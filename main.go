package main

import (
	 "github.com/rotblauer/trackpoints/trackPoint"
	 geo "github.com/rotblauer/google-geolocate"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"net/http"
	"flag"
	"log"

	"fmt"
	"time"
	"net"
)

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
	point, err:= client.Geolocate()
	if(err!=nil){
		fmt.Print(err)
	}
	fmt.Println(point)

	fmt.Printf("posting to   %s\n", postUrl)
	t := trackPoint.TrackPoint{Lat:point.Lat, Lng:point.Lng,Time:time.Now()}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(t)
	result, err := http.Post(postUrl, "application/json; charset=utf-8", b)
	if(err!=nil){
		fmt.Print(err)
	}
	io.Copy(os.Stdout, result.Body)

}

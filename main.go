package main

import (
	"flag"
	"fmt"
	"github.com/chr1sto14/dilgo/formathipchat"
	"github.com/chr1sto14/dilgo/net"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
	dilberturl := "http://dilbert.com"

	// get command line args
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [HIPCHAT-ROOM-URL]\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	url := flag.String("url", "", "url for hipchat room")

	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()

	if *url == "" {
		log.Fatalln("url is required")
	}

	log.Println("Fetching latest comic...")
	webbody, err := net.FetchUrl(dilberturl)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	msg, err := formathipchat.Format(webbody)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	err = net.PostMsg(*url, msg)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	log.Println("Comic posted to hipchat.")
}

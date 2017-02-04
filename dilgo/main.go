package main

import (
	"flag"
	"fmt"
	//"github/chr1sto14/dilbert/formathipchat"
	"github/chr1sto14/dilbert/net"
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
	webbytes, err := net.FetchUrl(dilberturl)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println("HTML:\n\n", string(webbytes))
	/*title, img, err =
	if err != nil {
		log.Fatalln("Error: ", err)
	}*/
}
